# ooe-impf-alert

Upper austria opened its registration for vaccination slots to the public and I wanted to find an earlier free slot than my current one for my district and the near vicinity.
So I hacked this little tool that reminds me whenever a free slot is available that is earlier than some specified date. At the moment this tool is simply running in the background and uses IFTTT to send push notifications to my phone so I can act instantaneously.

Thanks to [@internetztube](https://github.com/internetztube) who gave me this idea with his [ooe-impft-dates-overview](https://github.com/internetztube/ooe-impft-dates-overview) application.

## How to run this?

### CLI Tool

Example CLI command:

```
ooe-impf-alert-linux-amd64 --authorities=-400017 --date=2021-06-30 --ifttt-event-name=slot_found
```

Explanations for all the parameters for this call are directly below!

### Configuration

#### Vaccination authority

Find your authorities of interest using the [API](https://e-gov.ooe.gv.at/at.gv.ooe.cip/services/api/covid/authorities?adminUnitId=1&birthdate=1990-01-01) provided by the upper austrian government. Note down the `orgUnitId`s from your authorities.

#### IFTTT

Create a free account at https://ifttt.com/ and create an applet starting with a webhook followed by any action you may wanna take.
Note down the event name you provided to the webhook.
Furthermore when you click on "Documentation" while being on the created applet you find your personal key that you need to provide in order to trigger the webhook.

##### Data provided to the webhook

`value1` - the date and time of the slot

#### Date to tigger the notification

Mostly this will be the date you have your current appointment at. E.g. I want to be notified as soon as any appointment is found that is earlier than this date.
