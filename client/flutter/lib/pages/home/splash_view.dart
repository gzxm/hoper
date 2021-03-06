
import 'package:app/pages/home/splash_conroller.dart';
import 'package:extended_image/extended_image.dart';
import 'package:flutter/material.dart';
import 'package:flutter_spinkit/flutter_spinkit.dart';
import 'package:get/get.dart';
import 'package:get/get_state_manager/src/simple/get_state.dart';

final splash = Splash();

class Splash extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: ExtendedImage.asset(
        "assets/splash/splash.png",
        alignment: Alignment.center,
        width: Get.width,
        height: Get.height,
        fit: BoxFit.fill,
        //cancelToken: cancellationToken,
      ),
      floatingActionButton: GetBuilder<SplashController>(
        builder:(controller){
          if (controller.countdown == 0) {
            return Text('初始化');
          }
          return GestureDetector(
            onTap: (){
              if (! controller.adCompleter.isCompleted)  controller.adCompleter.complete();
            },
            child: Text('${controller.countdown}秒'),
          );
        }
      ),
    );
  }
}