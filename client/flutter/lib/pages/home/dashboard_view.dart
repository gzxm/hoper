import 'package:flutter/material.dart';

import 'package:get/get.dart';

import 'home_controller.dart';

class DashboardView extends GetView<HomeController> {
  @override
  Widget build(BuildContext context) {
    Get.log('DashboardView重绘');
    return Scaffold(
      body: Center(
        child: Column(
          mainAxisSize: MainAxisSize.min,
          children: [
            Text(
              'DashboardView is working',
              style: TextStyle(fontSize: 20),
            ),
            Obx(() => Text('Time: ${controller.now.toString()}',style: TextStyle(fontSize: 20),)),
          ],
        ),
      ),
    );
  }
}
