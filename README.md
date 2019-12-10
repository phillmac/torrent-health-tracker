# Torrent Health Tracker
![loc](https://tokei.rs/b1/github/nektro/torrent-health-tracker)
[![license](https://img.shields.io/github/license/nektro/torrent-health-tracker.svg)](https://github.com/nektro/torrent-health-tracker/blob/master/LICENSE)
[![astheno discord](https://img.shields.io/discord/551971034593755159.svg)](https://discord.gg/P6Y4zQC)
[![paypal](https://img.shields.io/badge/donate-paypal-009cdf)](https://paypal.me/nektro)
[![goreportcard](https://goreportcard.com/badge/github.com/nektro/torrent-health-tracker)](https://goreportcard.com/report/github.com/nektro/torrent-health-tracker)

**Torrent Health Tracker** is a way to provide a web interface into the statistics on the health of a collection of torrents. This is accomplished by looking at assembling tracker stats and periodically re-pinging the swarm. Simplely provide `torrent-health-tracker` with a list of trackers and magnet links and it'll do its thing!

> Note: currently only supports HTTP trackers.
