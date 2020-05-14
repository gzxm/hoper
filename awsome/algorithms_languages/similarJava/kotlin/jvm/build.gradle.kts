import org.jetbrains.kotlin.gradle.tasks.KotlinCompile

plugins {
    java
    kotlin("jvm") version "1.3.72"
}

group = "xyz.hoper"
version = "1.0-SNAPSHOT"

repositories {
    mavenCentral()
}

var junitJupiterEngineVersion = "5.4.0"

dependencies {
    implementation(kotlin("stdlib-jdk8"))
    implementation("org.objenesis:objenesis:3.0.1")
    implementation("org.apache.commons:commons-lang3:3.8.1")
    compileOnly("org.projectlombok:lombok:1.18.12")
    annotationProcessor("org.projectlombok:lombok:1.18.12")
    implementation("io.netty:netty-all:[5.0.0.Alpha2,)")
    testRuntimeOnly("org.junit.jupiter:junit-jupiter-engine:$junitJupiterEngineVersion")
    testImplementation("org.junit.jupiter:junit-jupiter-api:$junitJupiterEngineVersion")
    implementation("org.jetbrains.kotlin:kotlin-script-runtime:1.3.61")
}

configure<JavaPluginConvention> {
    sourceCompatibility = JavaVersion.VERSION_11
}

var moduleName = "jvm"


tasks {
    compileJava{
        options.encoding = "UTF-8"
        options.compilerArgs = listOf(
          "-Xlint:deprecation",
          "--add-opens=java.base/jdk.internal.misc=jvm",
          "--add-exports=java.base/jdk.internal.misc=jvm"
        )

    }

    compileKotlin {
        kotlinOptions.jvmTarget = "11"
        kotlinOptions.freeCompilerArgs = listOf("-Xinline-classes")
    }
    compileTestKotlin {
        kotlinOptions.jvmTarget = "11"
    }
}
