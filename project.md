What's the mvp?

People can create (very bare) profiles.
People can create.
People can contribute.
It's reasonably fluid to use.
There's no errors.

## Things to skimp on
- torrent system
- full profiles
- teams
- discovery system
- side products


## Full Auth Flow
- X login form
- X create account form
- X token on all authenticated requests
- redirect on token expiration
- "remember me"

## Home Page
- algorithm for shown
-- consider not having any, for fairness, instead have exploration and "subscription" systems where people can set notifications for things they're interested in. home page would focus on mission stuff and news
- header
-- link to login/profile
-- search
-- start new
-- explore

## Creation
- media
-- title
-- description
-- story
-- tags/category
-- location
- deliverables
- goal
- blind period
- promises
- side products
- finalize your account necessaries
- advice

## Creation Pages
- all media
-- video
-- demo materials
-- description
-- promises
- tickers
- websocket system for nearly complete?

## Profile Management
- create a profile
-- emails
-- display name
-- picture
-- description
-- links
-- vanity url
- account/security
-- notifications
-- log out on other
-- change password
-- payment methods

## Creation Discovery
- easy to search and explore
-- date started/funded/etc
-- current status
-- percentage funded
-- age
-- tags/category
-- location
-- all above with search
-- curated?
- easy to register specific notifications for *types* of things they're interested in hearing about

## Deployment
- netlify setup for client
- docker multistage build for server
- deployment script
- logging/notifications
- cd/ci system for server?
- nginx?

## Scaling
- cdn cache policies
- caching load balancer
- cache headers middleware
- autoscaling with supergiant
- prerendering of client
- regularly prerendering certain pages with actual api results

## File Upload and CDN
- X get files from client
- X upload to digital ocean space
- use keys for renders

## Full Security
- secure login and create account pages and forms
- secure card storage
- secure card page and form
- tls

## Payments and Escrow
- find provider
- secure credit card entry form
- remembering credit cards for accounts
- payment flow
- escrow mechanics
-- charge
-- release to creator

## Static Content
- mission
- explanation of mechanics
- beliefs

## Aesthetics and Skin
- great theme
- unified ux thinking

## Torrent Seeding System (future)
- create torrent files based on project
- gittorrent or other git system
- search engine for existing fulfilled thought work
- servers to act as guaranteed seeders
-- program that will actually do the seeding

## Teams (future)
- create a team
- add other people
- manage permissions

## Publicity Partnership (future)
- reach out to truly excellent and offer help getting the word out
