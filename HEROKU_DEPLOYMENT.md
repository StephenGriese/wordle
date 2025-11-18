# Deploying Wordle Helper to Heroku

## Overview
Complete guide to deploying the Wordle Helper web application to Heroku.

## Prerequisites

✅ Heroku account (you have one)
✅ Heroku CLI installed
✅ Git repository initialized
✅ Application committed to git

## Quick Deploy (5 Minutes)

### 1. Install Heroku CLI (if not already installed)
```bash
# macOS
brew install heroku/brew/heroku

# Or download from https://devcenter.heroku.com/articles/heroku-cli
```

### 2. Login to Heroku
```bash
heroku login
# Opens browser for authentication
```

### 3. Create Heroku App
```bash
cd /Users/sgries174@cable.comcast.com/repos/sjg/wordle

# Create app with auto-generated name
heroku create

# Or with custom name
heroku create your-wordle-helper

# Note the app URL: https://your-wordle-helper.herokuapp.com
```

### 4. Set Required Environment Variables
```bash
# Set dictionary path (relative to app root)
heroku config:set WORDLE_DICTIONARY=./american-english

# Set words to remove (optional)
heroku config:set WORDLE_REMOVE=./words-to-remove

# Verify configuration
heroku config
```

### 5. Set Go Buildpack
```bash
heroku buildpacks:set heroku/go
```

### 6. Deploy
```bash
# Add files if not already committed
git add Procfile go.mod cmd/ handlers/ dictionary/ views/ web/ wordle/ usrcmd/ scan/ pastwords/ american-english words-to-remove nytimes

# Commit
git commit -m "Add Heroku deployment configuration"

# Push to Heroku
git push heroku main

# Or if your branch is named differently
git push heroku master
```

### 7. Open Your App
```bash
heroku open
```

That's it! Your Wordle Helper is now live! 🎉

## Files Added for Heroku

### 1. Procfile
Tells Heroku how to run the app:
```
web: ./wordle-server
```

### 2. go.mod (if not exists)
Go module definition with dependencies.

### 3. Server Code Changes
- ✅ Reads `PORT` environment variable (Heroku requirement)
- ✅ Binds to `0.0.0.0` instead of `localhost`
- ✅ Falls back to `WORDLE_PORT` for local development

## Build Process

When you push to Heroku:

```
git push heroku main
    ↓
Heroku detects Go application
    ↓
Installs Go 1.22
    ↓
Runs: go build -o wordle-server cmd/server/main.go
    ↓
Creates slug with binary and static files
    ↓
Starts with: ./wordle-server
    ↓
App available on web
```

## Environment Variables

### Required
```bash
WORDLE_DICTIONARY=./american-english
```

### Optional
```bash
WORDLE_REMOVE=./words-to-remove  # Custom word exclusions
```

### Automatic (Set by Heroku)
```bash
PORT=xxxxx  # Dynamic port assigned by Heroku
```

## Heroku Commands Reference

### View Logs
```bash
# Tail logs in real-time
heroku logs --tail

# View recent logs
heroku logs --num 100

# Filter for errors
heroku logs --tail | grep ERROR
```

### Restart App
```bash
heroku restart
```

### Check Status
```bash
heroku ps
```

### Open App
```bash
heroku open
```

### Set Config
```bash
# Set variable
heroku config:set KEY=value

# View all config
heroku config

# Unset variable
heroku config:unset KEY
```

### Scale Dynos
```bash
# Check current dynos
heroku ps

# Scale up (for production)
heroku ps:scale web=2

# Scale down
heroku ps:scale web=1
```

### Run Commands
```bash
# Run arbitrary command
heroku run bash

# Check Go version
heroku run go version
```

## Deployment Workflow

### Initial Deployment
```bash
heroku create your-app-name
heroku config:set WORDLE_DICTIONARY=./american-english
git push heroku main
heroku open
```

### Subsequent Updates
```bash
# Make changes to code
git add .
git commit -m "Your changes"
git push heroku main

# Logs will show build and deployment
heroku logs --tail
```

### Rollback (if issues)
```bash
# View releases
heroku releases

# Rollback to previous version
heroku rollback v42
```

## Custom Domain (Optional)

### Add Custom Domain
```bash
# Add domain
heroku domains:add www.your-domain.com

# View DNS target
heroku domains
```

### Configure DNS
At your DNS provider, add CNAME record:
```
www.your-domain.com  →  your-app-name.herokuapp.com
```

### Enable SSL (Free)
```bash
heroku certs:auto:enable
```

## Performance Optimization

### Use Proper Dyno Type
```bash
# Free tier (sleeps after 30 min)
heroku ps:scale web=1:free

# Hobby tier ($7/mo, no sleep)
heroku ps:scale web=1:hobby

# Professional tier (better performance)
heroku ps:scale web=1:professional-1x
```

### Add New Relic Monitoring (Optional)
```bash
heroku addons:create newrelic:wayne
```

## Cost

### Free Tier
- ✅ Perfect for this app
- ✅ 550-1000 free dyno hours/month
- ⚠️ Sleeps after 30 minutes of inactivity
- ⚠️ Wake-up time: ~5-10 seconds

### Hobby Tier ($7/month)
- ✅ Never sleeps
- ✅ Custom domains with SSL
- ✅ Better for production use

### Professional Tier ($25-500/month)
- ✅ Horizontal scaling
- ✅ Metrics and monitoring
- ✅ High availability

**Recommendation:** Start with Free tier, upgrade to Hobby if you use it regularly.

## Troubleshooting

### Build Fails
```bash
# View build log
heroku logs --tail

# Common issues:
# 1. Missing go.mod → Create it
# 2. Import path errors → Fix module name
# 3. Missing files → Check .gitignore
```

### App Crashes
```bash
# Check logs
heroku logs --tail

# Check processes
heroku ps

# Restart app
heroku restart

# Run in debug mode
heroku config:set LOG_LEVEL=debug
```

### Environment Variables Not Set
```bash
# Check current config
heroku config

# Set missing variables
heroku config:set WORDLE_DICTIONARY=./american-english
```

### Port Binding Error
**Error:** `Error: server error: listen tcp 127.0.0.1:8080: bind: permission denied`

**Cause:** Not using Heroku's PORT environment variable

**Solution:** Already fixed! Server now reads `PORT` env var.

### Dictionary Not Found
**Error:** `failed to load dictionary: open ./american-english: no such file or directory`

**Cause:** File not committed to git

**Solution:**
```bash
git add american-english words-to-remove
git commit -m "Add dictionary files"
git push heroku main
```

### App Sleeps (Free Tier)
**Issue:** App slow to respond first time

**Explanation:** Free tier dynos sleep after 30 min inactivity

**Solutions:**
1. Upgrade to Hobby tier ($7/mo)
2. Use external pinger service (e.g., UptimeRobot)
3. Accept 5-10 second wake-up time

## Monitoring

### View Metrics
```bash
# View app metrics in browser
heroku dashboard

# Or command line
heroku ps
heroku logs --tail
```

### Set Up Alerts (Optional)
Use Heroku's metrics or integrate with:
- New Relic
- Datadog
- Sentry

## Security

### Environment Variables
✅ Never commit sensitive data
✅ Use Heroku config vars
✅ Dictionary files are public (OK for this app)

### HTTPS
✅ Heroku provides free SSL
✅ Automatic certificate management
✅ Force HTTPS (optional):
```go
// Add middleware in main.go
func redirectHTTPS(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.Header.Get("X-Forwarded-Proto") != "https" {
            http.Redirect(w, r, "https://"+r.Host+r.URL.String(), 
                http.StatusMovedPermanently)
            return
        }
        next.ServeHTTP(w, r)
    })
}
```

## Updating

### Deploy New Version
```bash
# Make changes
git add .
git commit -m "New feature"
git push heroku main

# Changes deploy automatically
```

### Zero-Downtime Deploys
Heroku provides preboot (requires Hobby+):
```bash
heroku features:enable preboot
```

## Backup and Maintenance

### Backup
No database, so just backup code:
```bash
git push origin main  # GitHub/GitLab
```

### Maintenance Mode
```bash
# Enable maintenance
heroku maintenance:on

# Disable maintenance
heroku maintenance:off
```

## Testing Deployed App

### Manual Test
```bash
# Open in browser
heroku open

# Or curl
curl https://your-app-name.herokuapp.com
```

### Reload Dictionary
```bash
# Test reload endpoint
curl -X POST https://your-app-name.herokuapp.com/reload
```

### Performance Test
```bash
# Simple load test
ab -n 100 -c 10 https://your-app-name.herokuapp.com/
```

## Alternative: Deploy to Other Platforms

### Fly.io
Similar to Heroku, good alternative:
```bash
fly launch
fly deploy
```

### Google Cloud Run
Serverless container deployment:
```bash
gcloud run deploy wordle-helper \
  --source . \
  --allow-unauthenticated
```

### Railway
Simple deployment:
```bash
railway login
railway init
railway up
```

## Summary

Deployment checklist:
- ✅ Procfile created
- ✅ go.mod exists
- ✅ Server uses PORT env var
- ✅ Server binds to 0.0.0.0
- ✅ Dictionary files committed
- ✅ Environment variables set
- ✅ Buildpack configured
- ✅ App deployed and tested

Your Wordle Helper is now live on Heroku! 🚀

## Quick Reference

```bash
# One-time setup
heroku login
heroku create your-app-name
heroku buildpacks:set heroku/go
heroku config:set WORDLE_DICTIONARY=./american-english

# Deploy
git push heroku main

# Manage
heroku open              # Open app
heroku logs --tail       # View logs
heroku restart           # Restart app
heroku config            # View config
heroku ps                # View dynos

# Update
git commit -am "changes"
git push heroku main
```

Enjoy your deployed Wordle Helper! 🎯

