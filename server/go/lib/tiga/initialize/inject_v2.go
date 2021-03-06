package initialize

import (
	"github.com/actliboy/hoper/server/go/lib/utils/log"
	"github.com/actliboy/hoper/server/go/lib/utils/reflect/mtos"
	"github.com/actliboy/hoper/server/go/lib/utils/slices"
	"github.com/pelletier/go-toml"
	"reflect"
	"strings"
)

func (init *Init) UnmarshalAndSetV2(bytes []byte) {
	tmp := map[string]interface{}{}
	err := toml.Unmarshal(bytes, &tmp)
	if err != nil {
		log.Fatal(err)
	}
	for k, v := range tmp {
		init.confM[strings.ToUpper(k)] = v
	}
	init.CloseDao()
	init.inject2()
}

// Customize
func (init *Init) inject2() {
	setConfig2(init.conf, init.confM)
	init.conf.Init()
	if init.dao == nil {
		return
	}
	setDao2(reflect.ValueOf(init.dao).Elem(), init.confM)
	init.dao.Init()
}

func setConfig2(conf Config, confM map[string]interface{}) {
	v := reflect.ValueOf(conf).Elem()
	for i := 0; i < v.NumField(); i++ {
		filed := v.Field(i)
		switch filed.Kind() {
		case reflect.Ptr:
			injectconf(filed.Interface(), strings.ToUpper(v.Type().Field(i).Name), confM)
		case reflect.Struct:
			injectconf(filed.Addr().Interface(), strings.ToUpper(v.Type().Field(i).Name), confM)
		}
	}
}

func injectconf(conf any, confName string, confM map[string]any) bool {
	filedv, ok := confM[confName]
	if ok {
		config := &mtos.DecoderConfig{
			Metadata: nil,
			Squash:   true,
			Result:   conf,
		}

		decoder, err := mtos.NewDecoder(config)
		if err != nil {
			return false
		}
		decoder.Decode(filedv)
	}
	return ok
}

func setDao2(v reflect.Value, confM map[string]any) {

	if !v.IsValid() {
		return
	}
	typ := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Addr().CanInterface() {
			if field.Kind() == reflect.Ptr && (!field.IsValid() || field.IsNil()) {
				field.Set(reflect.New(field.Type().Elem()))
			}
			inter := field.Addr().Interface()
			confName := strings.ToUpper(typ.Field(i).Name)
			if slices.StringContains(InitConfig.ConfigCenterConfig.NoInject, confName) {
				continue
			}
			if daofield, ok := inter.(DaoField); ok {
				tagSettings := ParseTagSetting(typ.Field(i).Tag.Get("init"), ";")
				if tagSettings.NotInject {
					continue
				}
				if tagSettings.ConfigName != "" {
					confName = tagSettings.ConfigName
				}
				conf := daofield.Config()
				/*
					??????conf???????????????????????????????????????????????????????????????????????????????????????????????????????????????interface{}
					valueConf := reflect.ValueOf(conf)
					if valueConf.Kind() == reflect.Ptr && (!valueConf.IsValid() || valueConf.IsNil()) {
						valueConf.Set(reflect.New(valueConf.Type().Elem()))
					}*/
				injectconf(conf, confName, confM)
				if conf1, ok := conf.(NeedInit); ok {
					conf1.Init()
				}
				if conf1, ok := conf.(Generate); ok {
					daofield.SetEntity(conf1.Generate())
				}
			}
		}
	}
}
