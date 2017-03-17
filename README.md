# Brighter, Blacker

Brighter, Blacker (BB) is a photo portfolio web application, meant to be used in conjunction with [bbclient](https://github.com/object88/bbclient) and  [bbservice](https://github.com/object88/bbservice).

## Architecture

On the client side, BB uses React to manage component rendering, and Relay to manage state & data fetching.

## Technical goals

* Use Google services to authenticate a user via SSO
* Transform relay's strange object shape into something more friendly towards react proptypes.

## Note
This application requires the schema file from `bbservice` to be copied to `/schema/schema.json`.

## Certificate creation

`openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout key.pem -out cert.pem`

#### Front end
* Go 1.8 serving JS bundles from bbclient
