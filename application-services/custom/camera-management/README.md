# Camera Management Example App Service
Use the Camera Management Example application service to auto discover and connect to nearby ONVIF and USB based cameras. The application will also control cameras via commands, create inference pipelines for the camera video streams and publish inference results to MQTT broker.

This app uses [EdgeX Core Services](https://github.com/edgexfoundry/edgex-compose.git), [Edgex Onvif device service](https://github.com/edgexfoundry/device-onvif-camera), [Edgex Usb device service](https://github.com/edgexfoundry/device-usb-camera) and [Edge Video Analytics Microservice](https://www.intel.com/content/www/us/en/developer/articles/technical/video-analytics-service.html).

A brief video demonstration of building and using the example app service can be found [here](https://www.youtube.com/watch?v=vZqd3j2Zn2Y).

## Steps for running this example:
Prerequisites:
A relatively modern Linux environment with `docker`, `docker-compose`, and `make` installed is required.


### 1. Start the EdgeX Core Services  

1. Clone `Edgex Compose` from github.com.
```shell 
git clone https://github.com/edgexfoundry/edgex-compose.git
``` 

1. Checkout `levski` branch.
```shell
git checkout levski
```

1. Navigate to the EdgeX `compose-builder` directory:

   ```bash
   cd edgex-compose/compose-builder/
   ```

1. update the `add-device-usb-camera.yml` file.:

   a. Add the rtsp server hostname environmental variable to the `device-usb-camera` service.
   ```yml
   services:
      device-usb-camera:
             environment:
               DRIVER_RTSPSERVERHOSTNAME: `your-local-ip-address`
   ```   

where `your-local-ip-address` is the ip address of the machine running the rtsp server.

   b. Under the `ports` section, find the one for port 8554 and change the host_ip from 127.0.0.1 to either 0.0.0.0 or the ip address you put in the previous step.

1. (Optional) If onvif cameras are being used, please refer to the [Edgex Onvif device service](https://github.com/edgexfoundry/device-onvif-camera/blob/main/doc/guides/CustomStartupGuide.md) documentation to run the onvif device service with camera credentials.

1. Run the following `make` command to run the edgex core services along with the Onvif and Usb device services.
```shell
make run no-secty ds-onvif-camera ds-usb-camera 
```   

### 2. Start [Edge Video Analytics Microservice](https://www.intel.com/content/www/us/en/developer/articles/technical/video-analytics-service.html) running for inference.

Note: Port for EVAM result streams has been changed from 8554 to 8555 to avoid conflicts with device-usb-camera service.

```shell
# Run this once to download edge-video-analytics into the edge-video-analytics sub-folder, 
# download models, and patch pipelines
make install-edge-video-analytics

# Run the EVAM services (in another terminal)
make run-edge-video-analytics
# ...
# OPTIONAL: Leave this running. If needed to stop
make stop-edge-video-analytics
```

4. Build and run the example application service. Web UI is used to view cameras, select models 
   and start inference pipelines for camera video streams and also view inference results streams.
```shell
# Build the app. 
make build-app

# Run the app.
make run-app
# ...
# Open your browser to http://localhost:59750
# ...
# Ctrl-C to stop it
```

### Development and Testing of UI
```shell
# Build the production web-ui into the web-ui/dist folder
# This is what is served by the app service on port 59750
make web-ui

# Serve the Web-UI in hot-reload mode on port 4200
make serve-ui
# ...
# Open your browser to http://localhost:4200
# ...
# Ctrl-C to stop it
```

[edgex-core-services]: https://github.com/edgexfoundry/edgex-go
[device-onvif-camera]: https://github.com/edgexfoundry-holding/device-onvif-camera
[evam]: https://www.intel.com/content/www/us/en/developer/articles/technical/video-analytics-service.html
