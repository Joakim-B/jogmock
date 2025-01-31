# Jogmock
A tool providing fake activity generation and uploads to online fitness trackers (currently only Strava is supported).

## Overview
Jogmock consumes routes (for example, constructed using [gpx.studio](https://gpx.studio/)) as input and automatically generates a realistic-looking activity in your favourite fitness tracker. It smooths out the provided route and generates the supposed speed on each point using a configurable random perlin-wave-based algorithm. The configuration provided can be used to tune it in order for the generated activities to look as close as possible to your real ones.

## Example
Using Jogmock is as simple as running a CLI app:
![Jogmock CLI Usage](./example/jogmock-cli.png)  

After which a new run magically appears in your activites!
![Jogmock result on Strava](./example/jogmock-strava.png)

## How to
1. Move config-example.yml to config.yml and change the variables to match your device. The following mobile versions can be used (or can be found in Mobile App -> Settings -> About) for the `mobile_app_version` version. Pick a random one, please.
   - 248.8 (1223885)
   - 247.10 (1223782)
   - 246.11 (1223676)
   - 245.9 (1223569)
   - 244.10 (1223463)
   - 243.9 (1223356)
   - 242.10 (1223251)
   - 241.8 (1223144)
   - 240.9 (1223041)
   - 238.9 (1223040)
   - 238.7 (1222828)
   - 237.5 (1222722)
   - 233.10 (1222306)
   - 233.9 (1222305)
   - 233.8 (1222304)
   - 231.9 (1222094)
   - 230.10 (1221988)
2. Use [gpx.studio](https://gpx.studio/) to draw your routes:
   1. Click `New GPX` at the top.
   2. Select Activity `run/hike` in the popup that appears on the left.
   3. Make sure that `Routing (follow roads)` is **NOT** selected.
   4. Click the minus on the right side of the popup to hide it.
   5. Click on the screen to add points, then export the file using Export (top of the screen) -> Download.
3. Using the modified config.yml and the downloaded file, you can start using Jogmock:
   1. Download the release appropriate to your OS and architecture.
   2. Install the bundled APK on your device. It is needed to get a reCAPTCHA token suitable for use with Strava and which will be needed during the first time you use `jogmock`.
   3. Launch the tool in your terminal and simply follow the steps. During GPX selection you can use Tab to autocomplete file paths.  

After configuring the reCAPTCHA token and login info, jogmock will cache it and not ask for it until needed.

## Help
The `run_activity` and `ride_activity` configs specify options used to generate the speed during an activity. You just need to specify the desired speed and `jogmock` will generate everything automatically, and these parameters can be used to fine-tune the generation in order for it to look as real (to you) as possible. Proper description of each option is specified in the example config.
