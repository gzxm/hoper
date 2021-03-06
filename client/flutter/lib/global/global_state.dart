
import 'dart:async';

import 'package:app/global/service.dart';

import 'package:flutter/cupertino.dart';

import '../pages/user/login_view.dart';
import 'app_info.dart';
import 'user.dart';
import 'package:get/get.dart';

import 'auth.dart';

export 'service.dart';

final globalState = GlobalState.instance;

class GlobalState extends GetxController{

  GlobalState._();

  static GlobalState? _instance;

  static GlobalState get instance => _instance ??= GlobalState._();

  var appState = AppInfo();
  var authState = AuthState();
  var userState = UserState();

  var initialized = false;

  Future<void> init() async {
    if (initialized) return;
    initialized = true;
    await globalService.init();
    await authState.getAuth();
  }

  Widget? authCheck() => authState.userAuth == null ? LoginView():null;

  var isDarkMode = (AppInfo.isDebug?true:false).obs;
}
