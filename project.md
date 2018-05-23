What's the mvp?

People can create (very bare) profiles.
People can create projects.
People can contribute to those projects.
It's reasonably fluid to use.
There's no errors.

Things to skimp on:
- torrent system
- full profiles
- teams
- discovery system
- side merch

[Full Auth Flow]
- login form
- create account form
- automatic renewal of tokens for active users
- "remember me"
- [token on all authenticated requests](https://github.com/axios/axios#creating-an-instance)

[Full Security]
- secure login and create account pages and forms
- secure card storage
- secure card page and form
- tls

[File Upload and CDN]
- get files from client
- upload to digital ocean space
- use keys for renders

[Payments and Escrow]
- find provider
- secure credit card entry form
- remembering credit cards for accounts
- payment flow
- escrow mechanics
-- charge
-- release to creator

[Static Content]
- mission
- explanation of mechanics
- beliefs

[Aesthetics and Skin]
- great theme
- unified component thoughts

[Home Page]
- algorithm for shown projects
-- consider not having any, for fairness, instead have exploration and "subscription" systems where people can set notifications for things they're interested in. home page would focus on mission stuff and news
- header
-- link to login/profile
-- search
-- start a new project
-- explore

[Project Creation]
- media
-- title
-- description
-- story
-- tags/category
-- location
- thought work deliverables
- goal
- blind period
- promises
- side merch
- finalize your account necessaries
- advice

[Project Pages]
- all media
-- video
-- demo materials
-- description
-- promises
- funding tickers
- websocket system for nearly complete projects?

[Profile Management]
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

[Overall Scaling]
- cdn cache policies
- caching responders
- auto-scaling systems for digital ocean

[Torrent Seeding System]
- create torrent files based on project
- gittorrent or other git system
- search engine for existing fulfilled thought work
- servers to act as guaranteed seeders
-- program that will actually do the seeding

[Project Discovery]
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

[Teams (future)]
- create a team
- add other people
- manage permissions

[Publicity Partnership (future)]
- reach out to truly excellent projects and offer help getting the word out


People need to be able to create projects, upload "demo" material of any kind, accept many potentially small payments that are charged immediately and held in escrow, and paid to them after the project is funded.

Once the project goes through, the information is released into the world with a permissive open culture license of some kind.
If it's code it could be released onto a git repo hub
If it's just files it could just be served on crowd sell
All files will be seeded for torrenting by crowd sell


The funding going through should be contingent on some sort of sanity check by the group


[Auth]
create new user
- X create users table
-- X migration
-- X unique constraint on email
- X receive email and password
- X create salt
- X hash password with salt
- X save both in database
- X sign a token
- (future?) send confirm email
login user
- X receive email and password
- X find user with email
-- X fail if not found
- X hash received password with salt
- X compare against database
-- X fail if not same
- X sign a token
