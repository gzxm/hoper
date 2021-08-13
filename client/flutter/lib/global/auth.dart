

import 'package:app/generated/protobuf/user/user.model.pb.dart';
import 'package:app/generated/protobuf/user/user.service.pb.dart';
import 'package:app/generated/protobuf/utils/empty/empty.pb.dart';
import 'package:app/global/service.dart';
import 'package:app/model/const/const.dart';
import 'package:app/pages/user/login_view.dart';
import 'package:app/utils/dialog.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:get/get_core/src/get_main.dart';
import 'package:get/get_navigation/src/extension_navigation.dart';
import 'package:grpc/grpc.dart';


class AuthState {
  UserAuthInfo? userAuth = null;
  User? self = null;

  static const _PRE = "AuthState";
  static const StringAuthKey = _PRE+Authorization;
  static const StringAccountKey = _PRE+"AccountKey";

  set account (String account)=> globalService.box.put(AuthState.StringAccountKey, account);
  String get account => globalService.box.get(AuthState.StringAccountKey);

  Future<void> getAuth() async {
    if (userAuth != null) return;
    final authKey = globalService.box.get(StringAuthKey);
    if (authKey != null) {
      try {
        final user = await globalService.userClient.stub.authInfo(Empty(),options:CallOptions(metadata: {Authorization: authKey}));
        this.userAuth = user;
        setAuth(authKey);
        return null;
      } catch (err) {
        print(err);
      }
    }
  }

  void setAuth(String authKey) {
    globalService.httpClient.options.headers[Authorization] = authKey;
    globalService.subject.setState(CallOptions(metadata: {Authorization: authKey}));
    globalService.box.put(AuthState.StringAuthKey, authKey);
  }

  Future<void> login(String account,String password) async{
    try{
      final rep = await globalService.userClient.stub.login(LoginReq(input: account, password: password,vCode: 'super'));
      final user = rep.user;
      self = rep.user;
      userAuth = UserAuthInfo(id:user.id,name:user.name,role:user.role,status:user.status);
      setAuth(rep.token);
      this.account = account;
      navigator!.pop();
    } on GrpcError catch (e) {
      dialog(e.message!);
    }catch (e) {
      // No specified type, handles all
      print('Something really unknown: $e');
    }
  }

  Future<void> logout() async{
    try{
      await globalService.userClient.stub.logout(Empty());
    } on GrpcError catch (e) {
      dialog(e.message!);
    }catch (e) {
      // No specified type, handles all
      print('Something really unknown: $e');
    }
  }
}


