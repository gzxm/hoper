import 'package:app/page/ffi/ffi.dart';
import 'package:app/page/lua/lua.dart';
import 'package:app/page/webview/webview.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:flutter_luakit_plugin/flutter_luakit_plugin.dart';
import 'page/bottom/bottom.dart';

void main() async {
  runApp(MaterialApp(home: MyApp()));
}

class MyApp extends StatelessWidget {
  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'hoper',
      theme: ThemeData(
        // This is the theme of your application.
        //
        // Try running your application with "flutter run". You'll see the
        // application has a blue toolbar. Then, without quitting the app, try
        // changing the primarySwatch below to Colors.green and then invoke
        // "hot reload" (press "r" in the console where you ran "flutter run",
        // or simply save your changes to "hot reload" in a Flutter IDE).
        // Notice that the counter didn't reset back to zero; the application
        // is not restarted.
        primarySwatch: Colors.blue,
        primaryColor: Colors.blueAccent[700],
        // This makes the visual density adapt to the platform that you run
        // the app on. For desktop platforms, the controls will be smaller and
        // closer together (more dense) than on mobile platforms.
        visualDensity: VisualDensity.adaptivePlatformDensity,
      ),
      home: MyHomePage(title: '🍀‬'),
    );
  }
}

class MyHomePage extends StatefulWidget {
  MyHomePage({Key key, this.title}) : super(key: key);

  // This widget is the home page of your application. It is stateful, meaning
  // that it has a State object (defined below) that contains fields that affect
  // how it looks.

  // This class is the configuration for the state. It holds the values (in this
  // case the title) provided by the parent (in this case the App widget) and
  // used by the build method of the State. Fields in a Widget subclass are
  // always marked "final".

  final String title;

  @override
  _MyHomePageState createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  final MethodChannel _methodChannel = const MethodChannel('xyz.hoper.native/view');
  int _counter = 0;
  int _selectedIndex = 0;
  static const TextStyle optionStyle =
      TextStyle(fontSize: 30, fontWeight: FontWeight.bold);

  void _onItemTapped(int index) {
    setState(() {
      _selectedIndex = index;
    });
  }

  void _incrementCounter() {
    setState(() {
      _counter++;
    });
  }

  void _reduceCounter() {
    setState(() {
      _counter--;
    });
  }

  AppBar bar(){
    return _selectedIndex!=1?AppBar(
      centerTitle: true,
      // Here we take the value from the MyHomePage object that was created by
      // the App.build method, and use it to set our appbar title.
      title: Text(widget.title),
    ):null;
  }

  @override
  Widget build(BuildContext context) {
    List<Widget> _widgetOptions = <Widget>[
      Column(
        mainAxisAlignment: MainAxisAlignment.center,
        children: <Widget>[
          Text(
            '点击下方加号:',
          ),
          Text(
            '$_counter',
            style: Theme.of(context).textTheme.headline4,
          ),
          Row(
              mainAxisAlignment: MainAxisAlignment.center,
              verticalDirection: VerticalDirection.up,
              children: [
                FloatingActionButton(
                  onPressed: _incrementCounter,
                  tooltip: 'Increment',
                  child: Icon(Icons.add),
                ),
                SizedBox(width: 20),
                FloatingActionButton(
                  onPressed: _reduceCounter,
                  tooltip: 'Reduce',
                  child: Icon(Icons.remove),
                )
              ]),
        ],
      ),
      WebViewExample(),
      Text(
        greeting(),
        style: optionStyle,
      ),
      Text(
        "lua",
        style: optionStyle,
      ),
    ];
    // This method is rerun every time setState is called, for instance as done
    // by the _incrementCounter method above.
    //
    // The Flutter framework has been optimized to make rerunning build methods
    // fast, so that you can just rebuild anything that needs updating rather
    // than having to individually change instances of widgets.
    return Scaffold(
      appBar: bar(),
      body: Center(
        // Center is a layout widget. It takes a single child and positions it
        // in the middle of the parent.
        child: _widgetOptions.elementAt(_selectedIndex),
      ),
      bottomNavigationBar: Bottom(onTap: _onItemTapped),
      floatingActionButton: FloatingActionButton(
        onPressed: ()=>_methodChannel.invokeMethod("toNative",{"route":"/"}).then((value) => null),
        tooltip: 'ToBrowser',
        child: Icon(Icons.send),
      ),
      // This trailing comma makes auto-formatting nicer for build methods.
      floatingActionButtonAnimator: FloatingActionButtonAnimator.scaling,
    );
  }
}
