FROM gradle:7.2.0-jdk17 as build

WORKDIR /github.com/s8sg/performance_test/spring_cloud_gw

COPY gradle             gradle
COPY src                src
COPY build.gradle       build.gradle
COPY gradlew            gradlew
COPY gradlew.bat        gradlew.bat
COPY settings.gradle    settings.gradle

RUN gradle check build --no-daemon --build-cache

FROM openjdk:17.0.1-jdk-slim as ship

COPY --from=build /github.com/s8sg/performance_test/spring_cloud_gw/build/libs/api-gateway-0.0.1-SNAPSHOT.jar .

EXPOSE 8080
CMD java -jar api-gateway-0.0.1-SNAPSHOT.jar
