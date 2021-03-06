import 'dart:async';


import 'package:app/components/weather/weather_bg.dart';
import 'package:app/components/weather/weather_rain_snow_bg.dart';

import 'package:app/components/weather/weather_type.dart';
import 'package:flutter/material.dart';
import 'package:get/get.dart';

class ChristmasTree extends StatefulWidget {
  ChristmasTree({Key? key}) : super(key: key);

  @override
  _ChristmasTreeState createState() => _ChristmasTreeState();
}

class _ChristmasTreeState extends State<ChristmasTree> with TickerProviderStateMixin {
  int branches = 7;
  late AnimationController controller;
  late Animation colorAnimation;

  @override
  void initState() {
    super.initState();
    controller =
        AnimationController(vsync: this, duration: Duration(milliseconds: 350));

    colorAnimation =
        ColorTween(begin: Colors.blue, end: Colors.red).animate(controller);
    controller.repeat(reverse: true);
    controller.addListener(()=>setState((){}));
  }

  @override
  void dispose() {
    controller.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Center(
        child: buildTree(),
    );
  }
  buildTree() {
    return Column(
      mainAxisAlignment: MainAxisAlignment.center,
      crossAxisAlignment: CrossAxisAlignment.center,
      children: [
        buildStar(),
        for (var i = 1; i < branches; i++) buildRow(i),
        buildBase(),
        buildBark(),
      ],
    );
  }

  buildRow(int i) {
    return Wrap(
      children: [
        for (var j = 0; j <= i; j++)
          Text(
            " * ",
            style: TextStyle(
              fontSize: 50,
              color: colorAnimation.value,
            ),
          ),
      ],
    );
  }

  Widget buildBase() {
    return Container(
      width: (branches-1).toDouble() * 50,
      height: 3.5,
      color: Colors.lightGreen,
    );
  }

  Widget buildBark() {
    return Container(
      width: 30,
      height: 100,
      color: Colors.green[700],
    );
  }
  Widget buildStar(){
    return Container(
      child: Icon(Icons.star,size:100,color: Colors.yellow,),
    );
  }
}

class SnowChristmasTree extends StatelessWidget{
  @override
  Widget build(BuildContext context) {
    return Scaffold(
        backgroundColor: Colors.black,
        body: SizeInherited(
            size: Size(Get.width, Get.height),
            child: Stack(
              children: [
                SizedBox(
                    width: Get.width,
                    height: Get.height,
                    child: ChristmasTree()),
                WeatherRainSnowBg(
                    weatherType: WeatherType.heavySnow,
                    viewWidth: Get.width,
                    viewHeight: Get.height),
              ],
            )));
  }

}