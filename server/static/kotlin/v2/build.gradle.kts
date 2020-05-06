import org.jetbrains.kotlin.gradle.tasks.KotlinCompile

plugins {
    id("org.springframework.boot") version "2.2.5.RELEASE"
    id("io.spring.dependency-management") version "1.0.9.RELEASE"
    kotlin("jvm") version "1.3.70"
    kotlin("plugin.spring") version "1.3.70"
}

ext{
    set("vertxVersion", "3.9.0")
    set("junitJupiterEngineVersion", "5.4.0")
    set("grpc_kotlin_version", "0.1.1")
    set("protobuf_version", "3.11.1")
    set("grpc_version", "1.28.0")
    set("kotlin_version", "1.3.70")
    set("springCloudAlibabaVersion", "2.2.0.RELEASE")
}

allprojects {
    apply<JavaPlugin>()
    group = "xyz.hoper"
    version = "0.0.1-SNAPSHOT"
    java.sourceCompatibility = JavaVersion.VERSION_11

    repositories {
        mavenCentral()
        gradlePluginPortal()
        google()
        jcenter()
        mavenLocal()
    }
}

subprojects{
    apply(plugin = "io.spring.dependency-management")
    apply(plugin = "org.springframework.boot")
    apply(plugin = "org.jetbrains.kotlin.jvm")
    apply(plugin = "org.jetbrains.kotlin.plugin.spring")


    val developmentOnly by configurations.creating
    configurations {
        runtimeClasspath {
            extendsFrom(developmentOnly)
        }
        compileOnly {
            extendsFrom(configurations.annotationProcessor.get())
        }
    }

    dependencies {
        implementation("org.springframework.boot:spring-boot-starter-webflux")
        implementation("org.apache.logging.log4j:log4j-core:2.12.1")
        implementation("com.fasterxml.jackson.module:jackson-module-kotlin")
        implementation("io.projectreactor.kotlin:reactor-kotlin-extensions")
        implementation("org.jetbrains.kotlin:kotlin-reflect")
        implementation("org.jetbrains.kotlin:kotlin-stdlib-jdk8")
        implementation("org.jetbrains.kotlinx:kotlinx-coroutines-reactor")
        compileOnly("org.projectlombok:lombok")
        developmentOnly("org.springframework.boot:spring-boot-devtools")
        annotationProcessor("org.projectlombok:lombok")
        testImplementation("org.springframework.boot:spring-boot-starter-test") {
            exclude(group = "org.junit.vintage", module = "junit-vintage-engine")
        }
        testImplementation("io.projectreactor:reactor-test")
        testImplementation("org.springframework.amqp:spring-rabbit-test")
        testImplementation("org.springframework.kafka:spring-kafka-test")
        testImplementation("org.springframework.security:spring-security-test")
        testImplementation("io.projectreactor:reactor-test")
    }

    dependencyManagement {
        imports {
            mavenBom("com.alibaba.cloud:spring-cloud-alibaba-dependencies:${property("springCloudAlibabaVersion")}")
        }
        dependencies {
            dependency("org.apache.logging.log4j:log4j-core:2.12.1")
        }
    }

    tasks.withType<Test> {
        useJUnitPlatform()
    }

    tasks.withType<KotlinCompile> {
        kotlinOptions {
            freeCompilerArgs = listOf("-Xjsr305=strict")
            jvmTarget = "1.8"
        }
    }

}
