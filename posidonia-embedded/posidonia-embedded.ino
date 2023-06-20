#include "deneyap.h"
#include <WiFi.h>
#include <WiFiAP.h>

#include "AsyncTCP.h"
#include "ESPAsyncWebServer.h"

#include <WiFiClient.h>

#define CAMERA_PIN 0             // i didnt read datasheet yet.

const char* ssid = "yigit";
const char* password = "sifrefalaniste";

AsyncWebServer server(80);

String header;

#include "esp_camera.h"
#include "SD_MMC.h"

// Select camera model
//#define CAMERA_MODEL_WROVER_KIT // Has PSRAM
//#define CAMERA_MODEL_ESP_EYE // Has PSRAM
//#define CAMERA_MODEL_M5STACK_PSRAM // Has PSRAM
//#define CAMERA_MODEL_M5STACK_V2_PSRAM // M5Camera version B Has PSRAM
//#define CAMERA_MODEL_M5STACK_WIDE // Has PSRAM
//#define CAMERA_MODEL_M5STACK_ESP32CAM // No PSRAM
#define CAMERA_MODEL_AI_THINKER // Has PSRAM
//#define CAMERA_MODEL_TTGO_T_JOURNAL // No PSRAM

#include "camera_pins.h"
#include "camera_config.h"

const char * photoPrefix = "/photo_";
int photoNumber = 0;

void setup() {
  Serial.begin(115200);
  Serial.setDebugOutput(true);
  Serial.println();



  if (psramFound()) {
    config.frame_size = FRAMESIZE_UXGA;
    config.jpeg_quality = 10;
    config.fb_count = 2;
  } else {
    config.frame_size = FRAMESIZE_SVGA;
    config.jpeg_quality = 12;
    config.fb_count = 1;
  }

  #if defined(CAMERA_MODEL_ESP_EYE)
  pinMode(13, INPUT_PULLUP);
  pinMode(14, INPUT_PULLUP);
  #endif

  // camera init
  esp_err_t err = esp_camera_init( & config);
  if (err != ESP_OK) {
    Serial.printf("Camera init failed with error 0x%x", err);
    return;
  }

  sensor_t * s = esp_camera_sensor_get();
  // initial sensors are flipped vertically and colors are a bit saturated
  if (s -> id.PID == OV3660_PID) {
    s -> set_vflip(s, 1); // flip it back
    s -> set_brightness(s, 1); // up the brightness just a bit
    s -> set_saturation(s, -2); // lower the saturation
  }
  // drop down frame size for higher initial frame rate
  s -> set_framesize(s, FRAMESIZE_QVGA);

  #if defined(CAMERA_MODEL_M5STACK_WIDE) || defined(CAMERA_MODEL_M5STACK_ESP32CAM)
  s -> set_vflip(s, 1);
  s -> set_hmirror(s, 1);
  #endif

  Serial.println("Initialising SD card");
  if (!SD_MMC.begin()) {
    Serial.println("Failed to initialise SD card!");
    return;
  }

  uint8_t cardType = SD_MMC.cardType();
  if (cardType == CARD_NONE) {
    Serial.println("SD card slot appears to be empty!");
    return;
  }

}

void loop() {
  camera_fb_t * fb = NULL;
  fb = esp_camera_fb_get();
  if (!fb) {
    Serial.println("Camera capture failed");
    return;
  }

  String photoFileName = photoPrefix + String(photoNumber) + ".jpg";
  fs::FS & fs = SD_MMC;
  Serial.printf("Picture file name: %s\n", photoFileName.c_str());

  File file = fs.open(photoFileName.c_str(), FILE_WRITE);
  if (!file) {
    Serial.println("Failed to open file in writing mode");
  } else {
    file.write(fb -> buf, fb -> len);
    Serial.printf("Saved file to path: %s\n", photoFileName.c_str());
    ++photoNumber;
  }
  file.close();
  esp_camera_fb_return(fb);

  delay(10000);
}