# API-Insights-UI

## How to run the Backend part that the UI part depends on it.
### 1. Enter the root folder of the git repo
### 2. Use the fllowing command in a terminal
```
docker-compose up mysql
```
### 3. Use the fllowing command in another terminal after step 2
```
docker-compose up backend
```
After that, the backed is running. and we can run UI part.

## How to run UI part
### 1. Enter /ui folder
### 2. Use the fllowing commands
```
npm install
npm start
```
Browser will open a new tab that its address is http://0.0.0.0:3000, the page will say "This site can’t be reached" at first and after a while it will works well without doing anything but wait.
