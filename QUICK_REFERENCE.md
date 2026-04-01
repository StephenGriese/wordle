# Quick Reference Card

## 🚀 Deploy to Heroku (Copy-Paste Ready)

```bash
# 1. Login
heroku login

# 2. Create app
heroku create my-wordle-helper

# 3. Configure
heroku buildpacks:set heroku/go
heroku config:set WORDLE_DICTIONARY=./american-english

# 4. Deploy
git add .
git commit -m "Update web app and deploy"
git push heroku main

# 5. Open
heroku open
```

**Done in 5 minutes!** ⏱️

---

## 📁 Files Ready to Commit

### New Features
- ✅ `dictionary/wordlist.go` - Thread-safe word manager
- ✅ `Procfile` - Heroku configuration
- ✅ `go.mod` - Go dependencies

### Modified
- ✅ `cmd/server/main.go` - Heroku + server startup
- ✅ `handlers/wordle.go` - Solve endpoint handlers
- ✅ `components/*.go` - Gomponents UI

### Documentation
- ✅ `SERVER_QUICKSTART.md`
- ✅ `HEROKU_DEPLOYMENT.md`
- ✅ `DOCUMENTATION_INDEX.md`
- ✅ Core architecture/design docs

---

## 📚 Dictionary Source

### Web + CLI Behavior
- Uses `WORDLE_DICTIONARY` (required)
- Optionally removes entries from `WORDLE_REMOVE`
- No past-word filtering by NYTimes list

---

## 🌐 Local vs Production

### Local Development
```bash
make run-server-dev
open http://localhost:8080
```

### Heroku Production
```bash
git push heroku main
heroku open
```

---

## 📊 Environment Variables

| Variable | Local | Heroku | Required |
|----------|-------|--------|----------|
| `WORDLE_DICTIONARY` | `./american-english` | Set via config | ✅ Yes |
| `WORDLE_REMOVE` | `./words-to-remove` | Set via config | ❌ No |
| `WORDLE_PORT` | `8080` | - | ❌ No |
| `PORT` | - | Auto-set | ✅ Yes (auto) |

---

## 🐛 Troubleshooting

### Build fails?
```bash
go mod tidy
go build cmd/server/main.go
```

### Heroku crashes?
```bash
heroku logs --tail
heroku restart
```

### Results look wrong?
1. Verify `WORDLE_DICTIONARY` points to the expected file
2. Verify `WORDLE_REMOVE` (if set) has expected entries
3. Check server logs

---

## 📖 Documentation

| Document | Purpose |
|----------|---------|
| `HEROKU_DEPLOYMENT.md` | 🚀 Deployment guide |
| `SERVER_QUICKSTART.md` | ⚡ Quick start guide |
| `DOCUMENTATION_INDEX.md` | 🧭 Complete documentation map |

---

## ✅ Pre-Commit Checklist

- [ ] Code compiles: `go build cmd/server/main.go`
- [ ] All files added: `git add .`
- [ ] Committed: `git commit -m "Update web app and deploy"`
- [ ] Dictionary files included
- [ ] go.mod exists
- [ ] Procfile exists

---

## 🎯 What You Built

### Features
1. ✅ **Web Solver** - Use the Wordle helper from browser
2. ✅ **Heroku Deploy** - Cloud hosting ready
3. ✅ **Thread Safety** - Concurrent access safe
4. ✅ **HTMX UI** - Smooth interactions
5. ✅ **Shared Logic** - CLI and web use the same solver

### Benefits
- 🚀 Deploy in 5 minutes
- 📚 Configurable dictionary source
- 🌐 Share with anyone
- 📱 Works on mobile
- 💰 Free tier available

---

## 💡 Pro Tips

### Free Heroku Tier
- Sleeps after 30 min
- 5-10 sec wake time
- Perfect for personal use

### Upgrade to Hobby ($7/mo)
```bash
heroku ps:scale web=1:hobby
```
Benefits: Never sleeps, faster, custom domain

### Monitor Your App
```bash
heroku logs --tail       # Watch logs
heroku ps                # Check status
heroku config            # View settings
```

---

## 🎉 Success!

Both features complete:
- ✅ Web solver working
- ✅ Heroku deployment ready
- ✅ Documentation complete
- ✅ Code compiles perfectly

**Next:** Test locally, commit, deploy! 🚀

