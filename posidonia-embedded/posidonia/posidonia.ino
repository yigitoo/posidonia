#include <WiFi.h>


int    HTTP_PORT   = 8080;
String HTTP_METHOD = "GET";
char   HOST_NAME[] = "example.phpoc.com";
String PATH_NAME   = "";

/*
* If you use ethernet you should declare with it
* EthernetClient client;
*/
WiFiClient client;

void setup() {
  Serial.begin(115200);
  if(client.connect(HOST_NAME, HTTP_PORT)) {
    Serial.println("Connected to server");
  } else {
    Serial.println("Connection failed!");
    for(;;){}
  }

  client.println(HTTP_METHOD + " " + PATH_NAME + "HTTP/1.1");
  client.println("Host: " + String(HOST_NAME));
  client.println("Connection: close");
  client.println();

}

void loop() {
  // put your main code here, to run repeatedly:

}
