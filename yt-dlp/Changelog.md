# Changelog

<!--
# To create a release, dispatch the https://github.com/yt-dlp/yt-dlp/actions/workflows/release.yml workflow on master
-->

### 2023.10.07

#### Extractor changes
- **abc.net.au**: iview: [Improve `episode` extraction](https://github.com/yt-dlp/yt-dlp/commit/a9efb4b8d74f3583450ffda0ee57259a47d39c70) ([#8201](https://github.com/yt-dlp/yt-dlp/issues/8201)) by [xofe](https://github.com/xofe)
- **erocast**: [Add extractor](https://github.com/yt-dlp/yt-dlp/commit/47c598783c98c179e04dd12c2a3fee0f3dc53087) ([#8264](https://github.com/yt-dlp/yt-dlp/issues/8264)) by [madewokherd](https://github.com/madewokherd)
- **gofile**: [Fix token cookie bug](https://github.com/yt-dlp/yt-dlp/commit/0730d5a966fa8a937d84bfb7f68be5198acb039b) by [bashonly](https://github.com/bashonly)
- **iq.com**: [Fix extraction and subtitles](https://github.com/yt-dlp/yt-dlp/commit/35d9cbaf9638ccc9daf8a863063b2e7c135bc664) ([#8260](https://github.com/yt-dlp/yt-dlp/issues/8260)) by [AS6939](https://github.com/AS6939)
- **lbry**
    - [Add playlist support](https://github.com/yt-dlp/yt-dlp/commit/48cceec1ddb8649b5e771df8df79eb9c39c82b90) ([#8213](https://github.com/yt-dlp/yt-dlp/issues/8213)) by [bashonly](https://github.com/bashonly), [drzraf](https://github.com/drzraf), [Grub4K](https://github.com/Grub4K)
    - [Extract `uploader_id`](https://github.com/yt-dlp/yt-dlp/commit/0e722f2f3ca42e634fd7b06ee70b16bf833ce132) ([#8244](https://github.com/yt-dlp/yt-dlp/issues/8244)) by [drzraf](https://github.com/drzraf)
- **litv**: [Fix extractor](https://github.com/yt-dlp/yt-dlp/commit/91a670a4f7babe9c8aa2018f57d8c8952a6f49d8) ([#7785](https://github.com/yt-dlp/yt-dlp/issues/7785)) by [jiru](https://github.com/jiru)
- **neteasemusic**: [Fix extractors](https://github.com/yt-dlp/yt-dlp/commit/f980df734cf5c0eaded2f7b38c6c60bccfeebb48) ([#8181](https://github.com/yt-dlp/yt-dlp/issues/8181)) by [c-basalt](https://github.com/c-basalt)
- **nhk**: [Fix VOD extraction](https://github.com/yt-dlp/yt-dlp/commit/e831c80e8b2fc025b3b67d82974cc59e3526fdc8) ([#8249](https://github.com/yt-dlp/yt-dlp/issues/8249)) by [garret1317](https://github.com/garret1317)
- **radiko**: [Improve extraction](https://github.com/yt-dlp/yt-dlp/commit/2ad3873f0dfa9285c91d2160e36c039e69d597c7) ([#8221](https://github.com/yt-dlp/yt-dlp/issues/8221)) by [garret1317](https://github.com/garret1317)
- **substack**
    - [Fix download cookies bug](https://github.com/yt-dlp/yt-dlp/commit/2f2dda3a7e85148773da3cdbc03ac9949ec1bc45) ([#8219](https://github.com/yt-dlp/yt-dlp/issues/8219)) by [handlerug](https://github.com/handlerug)
    - [Fix embed extraction](https://github.com/yt-dlp/yt-dlp/commit/fbcc299bd8a19cf8b3c8805d6c268a9110230973) ([#8218](https://github.com/yt-dlp/yt-dlp/issues/8218)) by [handlerug](https://github.com/handlerug)
- **theta**: [Remove extractors](https://github.com/yt-dlp/yt-dlp/commit/792f1e64f6a2beac51e85408d142b3118115c4fd) ([#8251](https://github.com/yt-dlp/yt-dlp/issues/8251)) by [alerikaisattera](https://github.com/alerikaisattera)
- **wrestleuniversevod**: [Call API with device ID](https://github.com/yt-dlp/yt-dlp/commit/b095fd3fa9d58a65dc9b830bd63b9d909422aa86) ([#8272](https://github.com/yt-dlp/yt-dlp/issues/8272)) by [bashonly](https://github.com/bashonly)
- **xhamster**: user: [Support creator urls](https://github.com/yt-dlp/yt-dlp/commit/cc8d8441524ec3442d7c0d3f8f33f15b66aa06f3) ([#8232](https://github.com/yt-dlp/yt-dlp/issues/8232)) by [Grub4K](https://github.com/Grub4K)
- **youtube**
    - [Fix `heatmap` extraction](https://github.com/yt-dlp/yt-dlp/commit/03e85ea99db76a2fddb65bf46f8819bda780aaf3) ([#8299](https://github.com/yt-dlp/yt-dlp/issues/8299)) by [bashonly](https://github.com/bashonly)
    - [Raise a warning for `Incomplete Data` instead of an error](https://github.com/yt-dlp/yt-dlp/commit/eb5bdbfa70126c7d5355cc0954b63720522e462c) ([#8238](https://github.com/yt-dlp/yt-dlp/issues/8238)) by [coletdjnz](https://github.com/coletdjnz)

#### Misc. changes
- **cleanup**
    - [Update extractor tests](https://github.com/yt-dlp/yt-dlp/commit/19c90e405b4137c06dfe6f9aaa02396df0da93e5) ([#7718](https://github.com/yt-dlp/yt-dlp/issues/7718)) by [trainman261](https://github.com/trainman261)
    - Miscellaneous: [377e85a](https://github.com/yt-dlp/yt-dlp/commit/377e85a1797db9e98b78b38203ed9d4ded229991) by [dirkf](https://github.com/dirkf), [gamer191](https://github.com/gamer191), [Grub4K](https://github.com/Grub4K)

### 2023.09.24

#### Important changes
- **The minimum *recommended* Python version has been raised to 3.8**
Since Python 3.7 has reached end-of-life, support for it will be dropped soon. [Read more](https://github.com/yt-dlp/yt-dlp/issues/7803)
- Security: [[CVE-2023-40581](https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2023-40581)] [Prevent RCE when using `--exec` with `%q` on Windows](https://github.com/yt-dlp/yt-dlp/security/advisories/GHSA-42h4-v29r-42qg)
    - The shell escape function is now using `""` instead of `\"`.
    - `utils.Popen` has been patched to properly quote commands.

#### Core changes
- [Fix HTTP headers and cookie handling](https://github.com/yt-dlp/yt-dlp/commit/6c5211cebeacfc53ad5d5ddf4a659be76039656f) by [bashonly](https://github.com/bashonly), [pukkandan](https://github.com/pukkandan)
- [Fix `--check-formats`](https://github.com/yt-dlp/yt-dlp/commit/8cb7fc44db010e965d808ee679ef0725cb6e147c) by [pukkandan](https://github.com/pukkandan)
- [Fix support for upcoming Python 3.12](https://github.com/yt-dlp/yt-dlp/commit/836e06d246512f286f30c1371b2c54b72c9ecd93) ([#8130](https://github.com/yt-dlp/yt-dlp/issues/8130)) by [Grub4K](https://github.com/Grub4K)
- [Merged with youtube-dl 66ab08](https://github.com/yt-dlp/yt-dlp/commit/9d6254069c75877bc88bc3584f4326fb1853a543) by [coletdjnz](https://github.com/coletdjnz)
- [Prevent RCE when using `--exec` with `%q` (CVE-2023-40581)](https://github.com/yt-dlp/yt-dlp/commit/de015e930747165dbb8fcd360f8775fd973b7d6e) by [Grub4K](https://github.com/Grub4K)
- [Raise minimum recommended Python version to 3.8](https://github.com/yt-dlp/yt-dlp/commit/61bdf15fc7400601c3da1aa7a43917310a5bf391) ([#8183](https://github.com/yt-dlp/yt-dlp/issues/8183)) by [Grub4K](https://github.com/Grub4K)
- [`FFmpegFixupM3u8PP` may need to run with ffmpeg](https://github.com/yt-dlp/yt-dlp/commit/f73c11803579889dc8e1c99e25dba9a22fef39d8) by [pukkandan](https://github.com/pukkandan)
- **compat**
    - [Add `types.NoneType`](https://github.com/yt-dlp/yt-dlp/commit/e0c4db04dc82a699bdabd9821ddc239ebe17d30a) by [pukkandan](https://github.com/pukkandan) (With fixes in [25b6e8f](https://github.com/yt-dlp/yt-dlp/commit/25b6e8f94679b4458550702b46e61249b875a4fd))
    - [Deprecate old functions](https://github.com/yt-dlp/yt-dlp/commit/3d2623a898196640f7cc0fc8b70118ff19e6925d) ([#2861](https://github.com/yt-dlp/yt-dlp/issues/2861)) by [coletdjnz](https://github.com/coletdjnz), [pukkandan](https://github.com/pukkandan)
    - [Ensure submodules are imported correctly](https://github.com/yt-dlp/yt-dlp/commit/a250b247334ce9f641e709cbb64974da6034a2b3) by [pukkandan](https://github.com/pukkandan)
- **cookies**: [Containers JSON should be opened as utf-8](https://github.com/yt-dlp/yt-dlp/commit/dab87ca23650fd87184ff5286b53e6985b59f71d) ([#7800](https://github.com/yt-dlp/yt-dlp/issues/7800)) by [bashonly](https://github.com/bashonly)
- **dependencies**: [Handle deprecation of `sqlite3.version`](https://github.com/yt-dlp/yt-dlp/commit/35f9a306e6934793cff100200cd03f288ec33f11) ([#8167](https://github.com/yt-dlp/yt-dlp/issues/8167)) by [bashonly](https://github.com/bashonly)
- **outtmpl**: [Fix replacement for `playlist_index`](https://github.com/yt-dlp/yt-dlp/commit/a264433c9fba147ecae2420091614186cfeeb895) by [pukkandan](https://github.com/pukkandan)
- **utils**
    - [Add temporary shim for logging](https://github.com/yt-dlp/yt-dlp/commit/1b392f905d20ef1f1b300b180f867d43c9ce49b8) by [pukkandan](https://github.com/pukkandan)
    - [Improve `parse_duration`](https://github.com/yt-dlp/yt-dlp/commit/af86873218c24c3859ccf575a87f2b00a73b49d0) by [bashonly](https://github.com/bashonly)
    - HTTPHeaderDict: [Handle byte values](https://github.com/yt-dlp/yt-dlp/commit/3f7965105d8d2048359e67c1e8b8ebd51588143b) by [pukkandan](https://github.com/pukkandan)
    - `clean_podcast_url`: [Handle more trackers](https://github.com/yt-dlp/yt-dlp/commit/2af4eeb77246b8183aae75a0a8d19f18c08115b2) ([#7556](https://github.com/yt-dlp/yt-dlp/issues/7556)) by [bashonly](https://github.com/bashonly), [mabdelfattah](https://github.com/mabdelfattah)
    - `js_to_json`: [Handle `Array` objects](https://github.com/yt-dlp/yt-dlp/commit/52414d64ca7b92d3f83964cdd68247989b0c4625) by [Grub4K](https://github.com/Grub4K), [std-move](https://github.com/std-move)

#### Extractor changes
- [Extract subtitles from SMIL manifests](https://github.com/yt-dlp/yt-dlp/commit/550e65410a7a1b105923494ac44460a4dc1a15d9) ([#7667](https://github.com/yt-dlp/yt-dlp/issues/7667)) by [bashonly](https://github.com/bashonly), [pukkandan](https://github.com/pukkandan)
- [Fix `--load-pages`](https://github.com/yt-dlp/yt-dlp/commit/81b4712bca608b9015aa68a4d96661d56e9cb894) by [pukkandan](https://github.com/pukkandan)
- [Make `_search_nuxt_data` more lenient](https://github.com/yt-dlp/yt-dlp/commit/904a19ee93195ce0bd4b08bd22b186120afb5b17) by [std-move](https://github.com/std-move)
- **abematv**
    - [Fix proxy handling](https://github.com/yt-dlp/yt-dlp/commit/497bbbbd7328cb705f70eced94dbd90993819a46) ([#8046](https://github.com/yt-dlp/yt-dlp/issues/8046)) by [SevenLives](https://github.com/SevenLives)
    - [Temporary fix for protocol handler](https://github.com/yt-dlp/yt-dlp/commit/9f66247289b9f8ecf931833b3f5f127274dd2161) by [pukkandan](https://github.com/pukkandan)
- **amazonminitv**: [Fix extractors](https://github.com/yt-dlp/yt-dlp/commit/538d37671a17e0782d17f08df17800e2e3bd57c8) by [bashonly](https://github.com/bashonly), [GautamMKGarg](https://github.com/GautamMKGarg)
- **antenna**: [Support antenna.gr](https://github.com/yt-dlp/yt-dlp/commit/665876034c8d3c031443f6b4958bed02ccdf4164) ([#7584](https://github.com/yt-dlp/yt-dlp/issues/7584)) by [stdedos](https://github.com/stdedos)
- **artetv**: [Fix HLS formats extraction](https://github.com/yt-dlp/yt-dlp/commit/c2da0b5ea215298135f76e3dc14b972a3c4afacb) by [bashonly](https://github.com/bashonly)
- **axs**: [Add extractor](https://github.com/yt-dlp/yt-dlp/commit/aee6b9b88c0bcccf27fd23b7e00fc0b7b168928f) ([#8094](https://github.com/yt-dlp/yt-dlp/issues/8094)) by [barsnick](https://github.com/barsnick)
- **banbye**: [Support video ids containing a hyphen](https://github.com/yt-dlp/yt-dlp/commit/578a82e497502b951036ce9da6fe0dac6937ac27) ([#8059](https://github.com/yt-dlp/yt-dlp/issues/8059)) by [kshitiz305](https://github.com/kshitiz305)
- **bbc**: [Extract tracklist as chapters](https://github.com/yt-dlp/yt-dlp/commit/eda0e415d26eb084e570cf5372d38ee1f616b70f) ([#7788](https://github.com/yt-dlp/yt-dlp/issues/7788)) by [garret1317](https://github.com/garret1317)
- **bild.de**: [Extract HLS formats](https://github.com/yt-dlp/yt-dlp/commit/b4c1c408c63724339eb12b16c91b253a7ee62cfa) ([#8032](https://github.com/yt-dlp/yt-dlp/issues/8032)) by [barsnick](https://github.com/barsnick)
- **bilibili**
    - [Add support for series, favorites and watch later](https://github.com/yt-dlp/yt-dlp/commit/9e68747f9607f05e92bb7d9b6e79d678b50070e1) ([#7518](https://github.com/yt-dlp/yt-dlp/issues/7518)) by [c-basalt](https://github.com/c-basalt)
    - [Extract Dolby audio formats](https://github.com/yt-dlp/yt-dlp/commit/b84fda7388dd20d38921e23b469147f3957c1812) ([#8142](https://github.com/yt-dlp/yt-dlp/issues/8142)) by [ClosedPort22](https://github.com/ClosedPort22)
    - [Extract `format_id`](https://github.com/yt-dlp/yt-dlp/commit/5336bf57a7061e0955a37f0542fc8ebf50d55b17) ([#7555](https://github.com/yt-dlp/yt-dlp/issues/7555)) by [c-basalt](https://github.com/c-basalt)
- **bilibilibangumi**: [Fix extractors](https://github.com/yt-dlp/yt-dlp/commit/bdd0b75e3f41ff35440eda6d395008beef19ef2f) ([#7337](https://github.com/yt-dlp/yt-dlp/issues/7337)) by [GD-Slime](https://github.com/GD-Slime)
- **bpb**: [Overhaul extractor](https://github.com/yt-dlp/yt-dlp/commit/f659e6439444ac64305b5c80688cd82f59d2279c) ([#8119](https://github.com/yt-dlp/yt-dlp/issues/8119)) by [Grub4K](https://github.com/Grub4K)
- **brilliantpala**: [Add extractors](https://github.com/yt-dlp/yt-dlp/commit/92feb5654c5a4c81ba872904a618700fcbb3e546) ([#6680](https://github.com/yt-dlp/yt-dlp/issues/6680)) by [pzhlkj6612](https://github.com/pzhlkj6612)
- **canal1, caracoltvplay**: [Add extractors](https://github.com/yt-dlp/yt-dlp/commit/b3febedbeb662dfdf9b5c1d5799039ad4fc969de) ([#7151](https://github.com/yt-dlp/yt-dlp/issues/7151)) by [elyse0](https://github.com/elyse0)
- **cbc**: [Ignore any 426 from API](https://github.com/yt-dlp/yt-dlp/commit/9bf14be775289bd88cc1f5c89fd761ae51879484) ([#7689](https://github.com/yt-dlp/yt-dlp/issues/7689)) by [makew0rld](https://github.com/makew0rld)
- **cbcplayer**: [Extract HLS formats and subtitles](https://github.com/yt-dlp/yt-dlp/commit/339c339fec095ff4141b20e6aa83629117fb26df) ([#7484](https://github.com/yt-dlp/yt-dlp/issues/7484)) by [trainman261](https://github.com/trainman261)
- **cbcplayerplaylist**: [Add extractor](https://github.com/yt-dlp/yt-dlp/commit/ed711897814f3ee0b1822e4205e74133467e8f1c) ([#7870](https://github.com/yt-dlp/yt-dlp/issues/7870)) by [trainman261](https://github.com/trainman261)
- **cineverse**: [Add extractors](https://github.com/yt-dlp/yt-dlp/commit/15591940ff102d1ae337d603a46d8f238c83a61f) ([#8146](https://github.com/yt-dlp/yt-dlp/issues/8146)) by [garret1317](https://github.com/garret1317)
- **crunchyroll**: [Remove initial state extraction](https://github.com/yt-dlp/yt-dlp/commit/9b16762f48914de9ac914601769c76668e433325) ([#7632](https://github.com/yt-dlp/yt-dlp/issues/7632)) by [Grub4K](https://github.com/Grub4K)
- **douyutv**: [Fix extractors](https://github.com/yt-dlp/yt-dlp/commit/21f40e75dfc0055ea9cdbd7fe2c46c6f9b561afd) ([#7652](https://github.com/yt-dlp/yt-dlp/issues/7652)) by [c-basalt](https://github.com/c-basalt)
- **dropbox**: [Fix extractor](https://github.com/yt-dlp/yt-dlp/commit/b9f2bc2dbed2323734a0d18e65e1e2e23dc833d8) ([#7926](https://github.com/yt-dlp/yt-dlp/issues/7926)) by [bashonly](https://github.com/bashonly), [denhotte](https://github.com/denhotte), [nathantouze](https://github.com/nathantouze) (With fixes in [099fb1b](https://github.com/yt-dlp/yt-dlp/commit/099fb1b35cf835303306549f5113d1802d79c9c7) by [bashonly](https://github.com/bashonly))
- **eplus**: inbound: [Add extractor](https://github.com/yt-dlp/yt-dlp/commit/295fbb3ae3a7d0dd50e286be5c487cf145ed5778) ([#5782](https://github.com/yt-dlp/yt-dlp/issues/5782)) by [pzhlkj6612](https://github.com/pzhlkj6612)
- **expressen**: [Improve `_VALID_URL`](https://github.com/yt-dlp/yt-dlp/commit/a5e264d74b4bd60c6e7ec4e38f1a23af4e420531) ([#8153](https://github.com/yt-dlp/yt-dlp/issues/8153)) by [kylegustavo](https://github.com/kylegustavo)
- **facebook**
    - [Add dash manifest URL](https://github.com/yt-dlp/yt-dlp/commit/a854fbec56d5004f5147116a41d1dd050632a579) ([#7743](https://github.com/yt-dlp/yt-dlp/issues/7743)) by [ringus1](https://github.com/ringus1)
    - [Fix webpage extraction](https://github.com/yt-dlp/yt-dlp/commit/d3d81cc98f554d0adb87d24bfd6fabaaa803944d) ([#7890](https://github.com/yt-dlp/yt-dlp/issues/7890)) by [ringus1](https://github.com/ringus1)
    - [Improve format sorting](https://github.com/yt-dlp/yt-dlp/commit/308936619c8a4f3a52d73c829c2006ff6c55fea2) ([#8074](https://github.com/yt-dlp/yt-dlp/issues/8074)) by [fireattack](https://github.com/fireattack)
    - reel: [Fix extraction](https://github.com/yt-dlp/yt-dlp/commit/bb5d84c9d2f1e978c3eddfb5ccbe138036682a36) ([#7564](https://github.com/yt-dlp/yt-dlp/issues/7564)) by [bashonly](https://github.com/bashonly), [demon071](https://github.com/demon071)
- **fox**: [Support foxsports.com](https://github.com/yt-dlp/yt-dlp/commit/30b29f37159e9226e2f2d5434c9a4096ac4efa2e) ([#7724](https://github.com/yt-dlp/yt-dlp/issues/7724)) by [ischmidt20](https://github.com/ischmidt20)
- **funker530**: [Fix extraction](https://github.com/yt-dlp/yt-dlp/commit/0ce1f48bf1cb78d40d734ce73ee1c90eccf92274) ([#8040](https://github.com/yt-dlp/yt-dlp/issues/8040)) by [04-pasha-04](https://github.com/04-pasha-04)
- **generic**
    - [Fix KVS thumbnail extraction](https://github.com/yt-dlp/yt-dlp/commit/53675852195d8dd859555d4789944a6887171ff8) by [bashonly](https://github.com/bashonly)
    - [Fix generic title for embeds](https://github.com/yt-dlp/yt-dlp/commit/994f7ef8e6003f4b7b258528755d0b6adcc31714) by [pukkandan](https://github.com/pukkandan)
- **gofile**: [Update token](https://github.com/yt-dlp/yt-dlp/commit/99c99c7185f5d8e9b3699a6fc7f86ec663d7b97e) by [bashonly](https://github.com/bashonly)
- **hotstar**
    - [Extract `release_year`](https://github.com/yt-dlp/yt-dlp/commit/7237c8dca0590aa7438ade93f927df88c9381ec7) ([#7869](https://github.com/yt-dlp/yt-dlp/issues/7869)) by [Rajeshwaran2001](https://github.com/Rajeshwaran2001)
    - [Make metadata extraction non-fatal](https://github.com/yt-dlp/yt-dlp/commit/30ea88591b728cca0896018dbf67c2298070c669) by [bashonly](https://github.com/bashonly)
    - [Support `/clips/` URLs](https://github.com/yt-dlp/yt-dlp/commit/86eeb044c2342d68c6ef177577f87852e6badd85) ([#7710](https://github.com/yt-dlp/yt-dlp/issues/7710)) by [bashonly](https://github.com/bashonly)
- **hungama**: [Overhaul extractors](https://github.com/yt-dlp/yt-dlp/commit/4b3a6ef1b3e235ba9a45142830b6edb357c71696) ([#7757](https://github.com/yt-dlp/yt-dlp/issues/7757)) by [bashonly](https://github.com/bashonly), [Yalab7](https://github.com/Yalab7)
- **indavideoembed**: [Fix extraction](https://github.com/yt-dlp/yt-dlp/commit/63e0c5748c0eb461a2ccca4181616eb930b4b750) ([#8129](https://github.com/yt-dlp/yt-dlp/issues/8129)) by [aky-01](https://github.com/aky-01)
- **iprima**: [Fix extractor](https://github.com/yt-dlp/yt-dlp/commit/568f08051841aedea968258889539741e26009e9) ([#7216](https://github.com/yt-dlp/yt-dlp/issues/7216)) by [std-move](https://github.com/std-move)
- **lbry**: [Fix original format extraction](https://github.com/yt-dlp/yt-dlp/commit/127a22460658ac39cbe5c4b3fb88d578363e0dfa) ([#7711](https://github.com/yt-dlp/yt-dlp/issues/7711)) by [bashonly](https://github.com/bashonly)
- **lecturio**: [Improve `_VALID_URL`](https://github.com/yt-dlp/yt-dlp/commit/efa2339502a37cf13ae7f143bd8b2c28f452d1cd) ([#7649](https://github.com/yt-dlp/yt-dlp/issues/7649)) by [simon300000](https://github.com/simon300000)
- **magellantv**: [Add extractor](https://github.com/yt-dlp/yt-dlp/commit/f4ea501551526ebcb54d19b84cf0ebe798583a85) ([#7616](https://github.com/yt-dlp/yt-dlp/issues/7616)) by [bashonly](https://github.com/bashonly)
- **massengeschmack.tv**: [Fix title extraction](https://github.com/yt-dlp/yt-dlp/commit/81f46ac573dc443ad48560f308582a26784d3015) ([#7813](https://github.com/yt-dlp/yt-dlp/issues/7813)) by [sb0stn](https://github.com/sb0stn)
- **media.ccc.de**: lists: [Fix extraction](https://github.com/yt-dlp/yt-dlp/commit/cf11b40ac40e3d23a6352753296f3a732886efb9) ([#8144](https://github.com/yt-dlp/yt-dlp/issues/8144)) by [Rohxn16](https://github.com/Rohxn16)
- **mediaite**: [Fix extraction](https://github.com/yt-dlp/yt-dlp/commit/630a55df8de7747e79aa680959d785dfff2c4b76) ([#7923](https://github.com/yt-dlp/yt-dlp/issues/7923)) by [Grabien](https://github.com/Grabien)
- **mediaklikk**: [Fix extractor](https://github.com/yt-dlp/yt-dlp/commit/6e07e4bc7e59f5bdb60e93c011e57b18b009f2b5) ([#8086](https://github.com/yt-dlp/yt-dlp/issues/8086)) by [bashonly](https://github.com/bashonly), [zhallgato](https://github.com/zhallgato)
- **mediastream**: [Make embed extraction non-fatal](https://github.com/yt-dlp/yt-dlp/commit/635ae31f68a3ac7f6393d59657ed711e34ee3552) by [bashonly](https://github.com/bashonly)
- **mixcloud**: [Update API URL](https://github.com/yt-dlp/yt-dlp/commit/7b71643cc986de9a3768dac4ac9b64f4d05e7f5e) ([#8114](https://github.com/yt-dlp/yt-dlp/issues/8114)) by [garret1317](https://github.com/garret1317)
- **monstercat**: [Add extractor](https://github.com/yt-dlp/yt-dlp/commit/eaee21bf71889d495076037cbe590c8c0b21ef3a) ([#8133](https://github.com/yt-dlp/yt-dlp/issues/8133)) by [garret1317](https://github.com/garret1317)
- **motortrendondemand**: [Update `_VALID_URL`](https://github.com/yt-dlp/yt-dlp/commit/c03a58ec9933e4a42c2d8fa80b8a0ddb2cde64e6) ([#7683](https://github.com/yt-dlp/yt-dlp/issues/7683)) by [AmirAflak](https://github.com/AmirAflak)
- **museai**: [Add extractor](https://github.com/yt-dlp/yt-dlp/commit/65cfa2b057d7946fbe322155a778fe206556d0c6) ([#7614](https://github.com/yt-dlp/yt-dlp/issues/7614)) by [bashonly](https://github.com/bashonly)
- **mzaalo**: [Improve `_VALID_URL`](https://github.com/yt-dlp/yt-dlp/commit/d7aee8e310b2c4f21d50aac0b420e1b3abde21a4) by [bashonly](https://github.com/bashonly)
- **n1info**: article: [Fix extractor](https://github.com/yt-dlp/yt-dlp/commit/8ac5b6d96ae5c60cd5ae2495949e0068a6754c45) ([#7373](https://github.com/yt-dlp/yt-dlp/issues/7373)) by [u-spec-png](https://github.com/u-spec-png)
- **nfl.com**: plus, replay: [Fix extractor](https://github.com/yt-dlp/yt-dlp/commit/1eaca74bc2ca0f5b1ec532f24c61de44f2e8cb2d) ([#7838](https://github.com/yt-dlp/yt-dlp/issues/7838)) by [bashonly](https://github.com/bashonly)
- **niconicochannelplus**: [Add extractors](https://github.com/yt-dlp/yt-dlp/commit/698beb9a497f51693e64d167e572ff9efa4bc25f) ([#5686](https://github.com/yt-dlp/yt-dlp/issues/5686)) by [pzhlkj6612](https://github.com/pzhlkj6612)
- **nitter**: [Fix title extraction fallback](https://github.com/yt-dlp/yt-dlp/commit/a83da3717d30697102e76f63a6f29d77f9373c2a) ([#8102](https://github.com/yt-dlp/yt-dlp/issues/8102)) by [ApoorvShah111](https://github.com/ApoorvShah111)
- **noodlemagazine**: [Fix extraction](https://github.com/yt-dlp/yt-dlp/commit/bae4834245a708fff97219849ec880c319c88bc6) ([#7830](https://github.com/yt-dlp/yt-dlp/issues/7830)) by [RedDeffender](https://github.com/RedDeffender) (With fixes in [69dbfe0](https://github.com/yt-dlp/yt-dlp/commit/69dbfe01c47cd078682a87f179f5846e2679e927) by [bashonly](https://github.com/bashonly))
- **novaembed**: [Fix extractor](https://github.com/yt-dlp/yt-dlp/commit/2269065ad60cb0ab62408ae6a7b20283e5252232) ([#7910](https://github.com/yt-dlp/yt-dlp/issues/7910)) by [std-move](https://github.com/std-move)
- **patreoncampaign**: [Fix extraction](https://github.com/yt-dlp/yt-dlp/commit/11de6fec9c9b8d34d1f90c8e6218ec58a3471b58) ([#7664](https://github.com/yt-dlp/yt-dlp/issues/7664)) by [bashonly](https://github.com/bashonly)
- **pbs**: [Add extractor `PBSKidsIE`](https://github.com/yt-dlp/yt-dlp/commit/6d6081dda1290a85bdab6717f239289e3aa74c8e) ([#7602](https://github.com/yt-dlp/yt-dlp/issues/7602)) by [snixon](https://github.com/snixon)
- **piapro**: [Support `/content` URL](https://github.com/yt-dlp/yt-dlp/commit/1bcb9fe8715b1f288efc322be3de409ee0597080) ([#7592](https://github.com/yt-dlp/yt-dlp/issues/7592)) by [FinnRG](https://github.com/FinnRG)
- **piaulizaportal**: [Add extractor](https://github.com/yt-dlp/yt-dlp/commit/6636021206dad17c7745ae6bce6cb73d6f2ef319) ([#7903](https://github.com/yt-dlp/yt-dlp/issues/7903)) by [pzhlkj6612](https://github.com/pzhlkj6612)
- **picartovod**: [Fix extractor](https://github.com/yt-dlp/yt-dlp/commit/db9743894071760f994f640a4c24358f749a78c0) ([#7727](https://github.com/yt-dlp/yt-dlp/issues/7727)) by [Frankgoji](https://github.com/Frankgoji)
- **pornbox**: [Add extractor](https://github.com/yt-dlp/yt-dlp/commit/40999467f72db074a3f13057da9bf82a857530fe) ([#7386](https://github.com/yt-dlp/yt-dlp/issues/7386)) by [niemands](https://github.com/niemands)
- **pornhub**: [Update access cookies for UK](https://github.com/yt-dlp/yt-dlp/commit/1d3d579c2142f69831b6ae140e1d8e824e07fa0e) ([#7591](https://github.com/yt-dlp/yt-dlp/issues/7591)) by [zhong-yiyu](https://github.com/zhong-yiyu)
- **pr0gramm**: [Rewrite extractor](https://github.com/yt-dlp/yt-dlp/commit/b532556d0a85e7d76f8f0880861232fb706ddbc5) ([#8151](https://github.com/yt-dlp/yt-dlp/issues/8151)) by [Grub4K](https://github.com/Grub4K)
- **radiofrance**: [Add support for livestreams, podcasts, playlists](https://github.com/yt-dlp/yt-dlp/commit/ba8e9eb2c8bbb699f314169fab8e544437ad731e) ([#7006](https://github.com/yt-dlp/yt-dlp/issues/7006)) by [elyse0](https://github.com/elyse0)
- **rbgtum**: [Fix extraction and support new URL format](https://github.com/yt-dlp/yt-dlp/commit/5fccabac27ca3c1165ade1b0df6fbadc24258dc2) ([#7690](https://github.com/yt-dlp/yt-dlp/issues/7690)) by [simon300000](https://github.com/simon300000)
- **reddit**
    - [Extract subtitles](https://github.com/yt-dlp/yt-dlp/commit/20c3c9b433dd47faf0dbde6b46e4e34eb76109a5) by [bashonly](https://github.com/bashonly)
    - [Fix thumbnail extraction](https://github.com/yt-dlp/yt-dlp/commit/9a04113dfbb69b904e4e2bea736da293505786b8) by [bashonly](https://github.com/bashonly)
- **rtvslo**: [Fix format extraction](https://github.com/yt-dlp/yt-dlp/commit/94389b225d9bcf29aa7ba8afaf1bbd7c62204eae) ([#8131](https://github.com/yt-dlp/yt-dlp/issues/8131)) by [bashonly](https://github.com/bashonly)
- **rule34video**: [Extract tags](https://github.com/yt-dlp/yt-dlp/commit/58493923e9b6f774947a2131e5258e9f3cf816be) ([#7117](https://github.com/yt-dlp/yt-dlp/issues/7117)) by [soundchaser128](https://github.com/soundchaser128)
- **rumble**: [Fix embed extraction](https://github.com/yt-dlp/yt-dlp/commit/23d829a3420450bcfb0788e6fb2cf4f6acdbe596) ([#8035](https://github.com/yt-dlp/yt-dlp/issues/8035)) by [trislee](https://github.com/trislee)
- **s4c**
    - [Add extractor](https://github.com/yt-dlp/yt-dlp/commit/b9de629d78ce31699f2de886071dc257830f9676) ([#7730](https://github.com/yt-dlp/yt-dlp/issues/7730)) by [ifan-t](https://github.com/ifan-t)
    - [Add series support and extract subs/thumbs](https://github.com/yt-dlp/yt-dlp/commit/fe371dcf0ba5ce8d42480eade54eeeac99ab3cb0) ([#7776](https://github.com/yt-dlp/yt-dlp/issues/7776)) by [ifan-t](https://github.com/ifan-t)
- **sohu**: [Fix extractor](https://github.com/yt-dlp/yt-dlp/commit/5be7e978867b5f66ad6786c674d79d40e950ae16) ([#7628](https://github.com/yt-dlp/yt-dlp/issues/7628)) by [bashonly](https://github.com/bashonly), [c-basalt](https://github.com/c-basalt)
- **stageplus**: [Fix m3u8 extraction](https://github.com/yt-dlp/yt-dlp/commit/56b3dc03354b75be995759d8441d2754c0442b9a) ([#7929](https://github.com/yt-dlp/yt-dlp/issues/7929)) by [bashonly](https://github.com/bashonly)
- **streamanity**: [Remove](https://github.com/yt-dlp/yt-dlp/commit/2cfe221fbbe46faa3f46552c08d947a51f424903) ([#7571](https://github.com/yt-dlp/yt-dlp/issues/7571)) by [alerikaisattera](https://github.com/alerikaisattera)
- **svtplay**: [Fix extraction](https://github.com/yt-dlp/yt-dlp/commit/2301b5c1b77a65abbb46b72f91e1e4666fd5d985) ([#7789](https://github.com/yt-dlp/yt-dlp/issues/7789)) by [dirkf](https://github.com/dirkf), [wader](https://github.com/wader)
- **tbsjp**: [Add episode, program, playlist extractors](https://github.com/yt-dlp/yt-dlp/commit/876b70c8edf4c0147f180bd981fbc4d625cbfb9c) ([#7765](https://github.com/yt-dlp/yt-dlp/issues/7765)) by [garret1317](https://github.com/garret1317)
- **tiktok**
    - [Fix audio-only format extraction](https://github.com/yt-dlp/yt-dlp/commit/b09bd0c19648f60c59fb980cd454cb0069959fb9) ([#7712](https://github.com/yt-dlp/yt-dlp/issues/7712)) by [bashonly](https://github.com/bashonly)
    - [Fix webpage extraction](https://github.com/yt-dlp/yt-dlp/commit/069cbece9dba6384f1cc5fcfc7ce562a31af42fc) by [bashonly](https://github.com/bashonly)
- **triller**: [Fix unlisted video extraction](https://github.com/yt-dlp/yt-dlp/commit/39837ae3199aa934299badbd0d63243ed639e6c8) ([#7670](https://github.com/yt-dlp/yt-dlp/issues/7670)) by [bashonly](https://github.com/bashonly)
- **tv5mondeplus**: [Fix extractor](https://github.com/yt-dlp/yt-dlp/commit/7d3d658f4c558ee7d72b1c01b46f2126948681cd) ([#7952](https://github.com/yt-dlp/yt-dlp/issues/7952)) by [dirkf](https://github.com/dirkf), [korli](https://github.com/korli)
- **twitcasting**
    - [Improve `_VALID_URL`](https://github.com/yt-dlp/yt-dlp/commit/cebbd33b1c678149fc8f0e254db6fc0da317ea80) ([#8120](https://github.com/yt-dlp/yt-dlp/issues/8120)) by [c-basalt](https://github.com/c-basalt)
    - [Support `--wait-for-video`](https://github.com/yt-dlp/yt-dlp/commit/c1d71d0d9f41db5e4306c86af232f5f6220a130b) ([#7975](https://github.com/yt-dlp/yt-dlp/issues/7975)) by [at-wat](https://github.com/at-wat)
- **twitter**
    - [Add fallback, improve error handling](https://github.com/yt-dlp/yt-dlp/commit/6014355c6142f68e20c8374e3787e5b5820f19e2) ([#7621](https://github.com/yt-dlp/yt-dlp/issues/7621)) by [bashonly](https://github.com/bashonly)
    - [Fix GraphQL and legacy API](https://github.com/yt-dlp/yt-dlp/commit/92315c03774cfabb3a921884326beb4b981f786b) ([#7516](https://github.com/yt-dlp/yt-dlp/issues/7516)) by [bashonly](https://github.com/bashonly)
    - [Fix retweet extraction and syndication API](https://github.com/yt-dlp/yt-dlp/commit/a006ce2b27357c15792eb5c18f06765e640b801c) ([#8016](https://github.com/yt-dlp/yt-dlp/issues/8016)) by [bashonly](https://github.com/bashonly)
    - [Revert 92315c03774cfabb3a921884326beb4b981f786b](https://github.com/yt-dlp/yt-dlp/commit/b03fa7834579a01cc5fba48c0e73488a16683d48) by [pukkandan](https://github.com/pukkandan)
    - spaces
        - [Fix format protocol](https://github.com/yt-dlp/yt-dlp/commit/613dbce177d34ffc31053e8e01acf4bb107bcd1e) ([#7550](https://github.com/yt-dlp/yt-dlp/issues/7550)) by [bashonly](https://github.com/bashonly)
        - [Pass referer header to downloader](https://github.com/yt-dlp/yt-dlp/commit/c6ef553792ed48462f9fd0e78143bef6b1a71c2e) by [bashonly](https://github.com/bashonly)
- **unsupported**: [List more sites with DRM](https://github.com/yt-dlp/yt-dlp/commit/e7057383380d7d53815f8feaf90ca3dcbde88983) by [pukkandan](https://github.com/pukkandan)
- **videa**: [Fix extraction](https://github.com/yt-dlp/yt-dlp/commit/98eac0e6ba0e510ae7dfdfd249d42ee71fb272b1) ([#8003](https://github.com/yt-dlp/yt-dlp/issues/8003)) by [aky-01](https://github.com/aky-01), [hatsomatt](https://github.com/hatsomatt)
- **vrt**: [Update token signing key](https://github.com/yt-dlp/yt-dlp/commit/325191d0c9bf3fe257b8a7c2eb95080f44f6ddfc) ([#7519](https://github.com/yt-dlp/yt-dlp/issues/7519)) by [Zprokkel](https://github.com/Zprokkel)
- **wat.tv**: [Fix extraction](https://github.com/yt-dlp/yt-dlp/commit/7cccab79e7d00ed965b48b8cefce1da8a0513409) ([#7898](https://github.com/yt-dlp/yt-dlp/issues/7898)) by [davinkevin](https://github.com/davinkevin)
- **wdr**: [Fix extraction](https://github.com/yt-dlp/yt-dlp/commit/5d0395498d7065aa5e55bac85fa9354b4b0d48eb) ([#7979](https://github.com/yt-dlp/yt-dlp/issues/7979)) by [szabyg](https://github.com/szabyg)
- **web.archive**: vlive: [Remove extractor](https://github.com/yt-dlp/yt-dlp/commit/9652bca1bd02f6bc1b8cb1e186f2ccbf32225561) ([#8132](https://github.com/yt-dlp/yt-dlp/issues/8132)) by [bashonly](https://github.com/bashonly)
- **weibo**: [Fix extractor and support user extraction](https://github.com/yt-dlp/yt-dlp/commit/69b03f84f8378b0b5a2fbae56f9b7d860b2f529e) ([#7657](https://github.com/yt-dlp/yt-dlp/issues/7657)) by [c-basalt](https://github.com/c-basalt)
- **weverse**: [Support extraction without auth](https://github.com/yt-dlp/yt-dlp/commit/c2d8ee0000302aba63476b7d5bd8793e57b6c8c6) ([#7924](https://github.com/yt-dlp/yt-dlp/issues/7924)) by [seproDev](https://github.com/seproDev)
- **wimbledon**: [Add extractor](https://github.com/yt-dlp/yt-dlp/commit/a15fcd299e767a510debd8dc1646fe863b96ce0e) ([#7551](https://github.com/yt-dlp/yt-dlp/issues/7551)) by [nnoboa](https://github.com/nnoboa)
- **wrestleuniverseppv**: [Fix HLS AES key extraction](https://github.com/yt-dlp/yt-dlp/commit/dae349da97cafe7357106a8f3187fd48a2ad1210) by [bashonly](https://github.com/bashonly)
- **youtube**
    - [Add `player_params` extractor arg](https://github.com/yt-dlp/yt-dlp/commit/ba06d77a316650ff057347d224b5afa8b203ad65) ([#7719](https://github.com/yt-dlp/yt-dlp/issues/7719)) by [coletdjnz](https://github.com/coletdjnz)
    - [Fix `player_params` arg being converted to lowercase](https://github.com/yt-dlp/yt-dlp/commit/546b2c28a106cf8101d481b215b676d1b091d276) by [coletdjnz](https://github.com/coletdjnz)
    - [Fix consent cookie](https://github.com/yt-dlp/yt-dlp/commit/378ae9f9fb8e8c86e6ac89c4c5b815b48ce93620) ([#7774](https://github.com/yt-dlp/yt-dlp/issues/7774)) by [coletdjnz](https://github.com/coletdjnz)
    - tab: [Detect looping feeds](https://github.com/yt-dlp/yt-dlp/commit/1ba6fe9db5f660d5538588315c23ad6cf0371c5f) ([#6621](https://github.com/yt-dlp/yt-dlp/issues/6621)) by [coletdjnz](https://github.com/coletdjnz)
- **zaiko**: [Improve thumbnail extraction](https://github.com/yt-dlp/yt-dlp/commit/ecef42c3adbcb6a84405139047923c4967316f28) ([#8054](https://github.com/yt-dlp/yt-dlp/issues/8054)) by [pzhlkj6612](https://github.com/pzhlkj6612)
- **zee5**: [Update access token endpoint](https://github.com/yt-dlp/yt-dlp/commit/a0de8bb8601146b8f87bf7cd562eef8bfb4690be) ([#7914](https://github.com/yt-dlp/yt-dlp/issues/7914)) by [bashonly](https://github.com/bashonly)
- **zoom**: [Extract duration](https://github.com/yt-dlp/yt-dlp/commit/66cc64ff6696f9921ff112a278542f8d999ffea4) by [bashonly](https://github.com/bashonly)

#### Downloader changes
- **external**
    - [Fix ffmpeg input from stdin](https://github.com/yt-dlp/yt-dlp/commit/e57eb98222d29cc4c09ee975d3c492274a6e5be3) ([#7655](https://github.com/yt-dlp/yt-dlp/issues/7655)) by [bashonly](https://github.com/bashonly)
    - [Fixes to cookie handling](https://github.com/yt-dlp/yt-dlp/commit/42ded0a429c20ec13dc006825e1508d9a02f0ad4) by [bashonly](https://github.com/bashonly)

#### Postprocessor changes
- **embedthumbnail**: [Support `m4v`](https://github.com/yt-dlp/yt-dlp/commit/8a4cd12c8f8e93292e3e95200b9d17a3af39624c) ([#7583](https://github.com/yt-dlp/yt-dlp/issues/7583)) by [Neurognostic](https://github.com/Neurognostic)

#### Networking changes
- [Add module](https://github.com/yt-dlp/yt-dlp/commit/c365dba8430ee33abda85d31f95128605bf240eb) ([#2861](https://github.com/yt-dlp/yt-dlp/issues/2861)) by [pukkandan](https://github.com/pukkandan)
- [Add request handler preference framework](https://github.com/yt-dlp/yt-dlp/commit/db7b054a6111ca387220d0eb87bf342f9c130eb8) ([#7603](https://github.com/yt-dlp/yt-dlp/issues/7603)) by [coletdjnz](https://github.com/coletdjnz)
- [Add strict Request extension checking](https://github.com/yt-dlp/yt-dlp/commit/86aea0d3a213da3be1da638b9b828e6f0ee1d59f) ([#7604](https://github.com/yt-dlp/yt-dlp/issues/7604)) by [coletdjnz](https://github.com/coletdjnz)
- [Fix POST requests with zero-length payloads](https://github.com/yt-dlp/yt-dlp/commit/71baa490ebd3655746430f208a9b605d120cd315) ([#7648](https://github.com/yt-dlp/yt-dlp/issues/7648)) by [bashonly](https://github.com/bashonly)
- [Fix `--legacy-server-connect`](https://github.com/yt-dlp/yt-dlp/commit/75dc8e673b481a82d0688aeec30f6c65d82bb359) ([#7645](https://github.com/yt-dlp/yt-dlp/issues/7645)) by [bashonly](https://github.com/bashonly)
- [Fix various socks proxy bugs](https://github.com/yt-dlp/yt-dlp/commit/20fbbd9249a2f26c7ae579bde5ba5d69aa8fac69) ([#8065](https://github.com/yt-dlp/yt-dlp/issues/8065)) by [coletdjnz](https://github.com/coletdjnz)
- [Ignore invalid proxies in env](https://github.com/yt-dlp/yt-dlp/commit/bbeacff7fcaa3b521066088a5ccbf34ef5070d1d) ([#7704](https://github.com/yt-dlp/yt-dlp/issues/7704)) by [coletdjnz](https://github.com/coletdjnz)
- [Rewrite architecture](https://github.com/yt-dlp/yt-dlp/commit/227bf1a33be7b89cd7d44ad046844c4ccba104f4) ([#2861](https://github.com/yt-dlp/yt-dlp/issues/2861)) by [coletdjnz](https://github.com/coletdjnz)
- **Request Handler**
    - urllib
        - [Remove dot segments during URL normalization](https://github.com/yt-dlp/yt-dlp/commit/4bf912282a34b58b6b35d8f7e6be535770c89c76) ([#7662](https://github.com/yt-dlp/yt-dlp/issues/7662)) by [coletdjnz](https://github.com/coletdjnz)
        - [Simplify gzip decoding](https://github.com/yt-dlp/yt-dlp/commit/59e92b1f1833440bb2190f847eb735cf0f90bc85) ([#7611](https://github.com/yt-dlp/yt-dlp/issues/7611)) by [Grub4K](https://github.com/Grub4K) (With fixes in [77bff23](https://github.com/yt-dlp/yt-dlp/commit/77bff23ee97565bab2e0d75b893a21bf7983219a))

#### Misc. changes
- **build**: [Make sure deprecated modules are added](https://github.com/yt-dlp/yt-dlp/commit/131d132da5c98c6c78bd7eed4b37f4458561b3d9) by [pukkandan](https://github.com/pukkandan)
- **cleanup**
    - [Add color to `download-archive` message](https://github.com/yt-dlp/yt-dlp/commit/2b029ca0a9f9105c4f7626993fa60e54c9782749) ([#5138](https://github.com/yt-dlp/yt-dlp/issues/5138)) by [aaruni96](https://github.com/aaruni96), [Grub4K](https://github.com/Grub4K), [pukkandan](https://github.com/pukkandan)
    - Miscellaneous
        - [6148833](https://github.com/yt-dlp/yt-dlp/commit/6148833f5ceb7674142ddb8d761ffe03cee7df69), [62b5c94](https://github.com/yt-dlp/yt-dlp/commit/62b5c94cadaa5f596dc1a7083db9db12efe357be) by [pukkandan](https://github.com/pukkandan)
        - [5ca095c](https://github.com/yt-dlp/yt-dlp/commit/5ca095cbcde3e32642a4fe5b2d69e8e3c785a021) by [barsnick](https://github.com/barsnick), [bashonly](https://github.com/bashonly), [coletdjnz](https://github.com/coletdjnz), [gamer191](https://github.com/gamer191), [Grub4K](https://github.com/Grub4K), [sqrtNOT](https://github.com/sqrtNOT)
        - [088add9](https://github.com/yt-dlp/yt-dlp/commit/088add9567d39b758737e4299a0e619fd89d2e8f) by [Grub4K](https://github.com/Grub4K)
- **devscripts**: `make_changelog`: [Fix changelog grouping and add networking group](https://github.com/yt-dlp/yt-dlp/commit/30ba233d4cee945756ed7344e7ddb3a90d2ae608) ([#8124](https://github.com/yt-dlp/yt-dlp/issues/8124)) by [Grub4K](https://github.com/Grub4K)
- **docs**: [Update collaborators](https://github.com/yt-dlp/yt-dlp/commit/1be0a96a4d14f629097509fcc89d15f69a8243c7) by [Grub4K](https://github.com/Grub4K)
- **test**
    - [Add tests for socks proxies](https://github.com/yt-dlp/yt-dlp/commit/fcd6a76adc49d5cd8783985c7ce35384b72e545f) ([#7908](https://github.com/yt-dlp/yt-dlp/issues/7908)) by [coletdjnz](https://github.com/coletdjnz)
    - [Fix `httplib_validation_errors` test for old Python versions](https://github.com/yt-dlp/yt-dlp/commit/95abea9a03289da1384e5bda3d590223ccc0a238) ([#7677](https://github.com/yt-dlp/yt-dlp/issues/7677)) by [coletdjnz](https://github.com/coletdjnz)
    - [Fix `test_load_certifi`](https://github.com/yt-dlp/yt-dlp/commit/de20687ee6b742646128a7629b57096631a20619) by [pukkandan](https://github.com/pukkandan)
    - download: [Test for `expected_exception`](https://github.com/yt-dlp/yt-dlp/commit/661c9a1d029296b28e0b2f8be8a72a43abaf6536) by [at-wat](https://github.com/at-wat)

### 2023.07.06

#### Important changes
- Security: [[CVE-2023-35934](https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2023-35934)] Fix [Cookie leak](https://github.com/yt-dlp/yt-dlp/security/advisories/GHSA-v8mc-9377-rwjj)
    - `--add-header Cookie:` is deprecated and auto-scoped to input URL domains
    - Cookies are scoped when passed to external downloaders
    - Add `cookies` field to info.json and deprecate `http_headers.Cookie`

#### Core changes
- [Allow extractors to mark formats as potentially DRM](https://github.com/yt-dlp/yt-dlp/commit/bc344cd456380999c1ee74554dfd432a38f32ec7) ([#7396](https://github.com/yt-dlp/yt-dlp/issues/7396)) by [pukkandan](https://github.com/pukkandan)
- [Bugfix for b4e0d75848e9447cee2cd3646ce54d4744a7ff56](https://github.com/yt-dlp/yt-dlp/commit/e59e20744eb32ce4b6ea0dece7c673be8376a710) by [pukkandan](https://github.com/pukkandan)
- [Change how `Cookie` headers are handled](https://github.com/yt-dlp/yt-dlp/commit/3121512228487c9c690d3d39bfd2579addf96e07) by [Grub4K](https://github.com/Grub4K)
- [Prevent `Cookie` leaks on HTTP redirect](https://github.com/yt-dlp/yt-dlp/commit/f8b4bcc0a791274223723488bfbfc23ea3276641) by [coletdjnz](https://github.com/coletdjnz)
- **formats**: [Fix best fallback for storyboards](https://github.com/yt-dlp/yt-dlp/commit/906c0bdcd8974340d619e99ccd613c163eb0d0c2) by [pukkandan](https://github.com/pukkandan)
- **outtmpl**: [Pad `playlist_index` etc even when with internal formatting](https://github.com/yt-dlp/yt-dlp/commit/47bcd437247152e0af5b3ebc5592db7bb66855c2) by [pukkandan](https://github.com/pukkandan)
- **utils**: clean_podcast_url: [Handle protocol in redirect URL](https://github.com/yt-dlp/yt-dlp/commit/91302ed349f34dc26cc1d661bb45a4b71f4417f7) by [pukkandan](https://github.com/pukkandan)

#### Extractor changes
- **abc**: [Fix extraction](https://github.com/yt-dlp/yt-dlp/commit/8f05fbae2a79ce0713077ccc68b354e63216bf20) ([#7434](https://github.com/yt-dlp/yt-dlp/issues/7434)) by [meliber](https://github.com/meliber)
- **AdultSwim**: [Extract subtitles from m3u8](https://github.com/yt-dlp/yt-dlp/commit/5e16cf92eb496b7c1541a6b1d727cb87542984db) ([#7421](https://github.com/yt-dlp/yt-dlp/issues/7421)) by [nnoboa](https://github.com/nnoboa)
- **crunchyroll**: music: [Fix `_VALID_URL`](https://github.com/yt-dlp/yt-dlp/commit/5b4b92769afcc398475e481bfa839f1158902fe9) ([#7439](https://github.com/yt-dlp/yt-dlp/issues/7439)) by [AmanSal1](https://github.com/AmanSal1), [rdamas](https://github.com/rdamas)
- **Douyin**: [Fix extraction from webpage](https://github.com/yt-dlp/yt-dlp/commit/a2be9781fbf4d7e4db245c277ca2ecc41cf3a7b2) by [bashonly](https://github.com/bashonly)
- **googledrive**: [Fix source format extraction](https://github.com/yt-dlp/yt-dlp/commit/3b7f5300c577fef40464d46d4e4037a69d51fe82) ([#7395](https://github.com/yt-dlp/yt-dlp/issues/7395)) by [RfadnjdExt](https://github.com/RfadnjdExt)
- **kick**: [Fix `_VALID_URL`](https://github.com/yt-dlp/yt-dlp/commit/ef8509c300ea50da86aea447eb214d3d6f6db6bb) by [bashonly](https://github.com/bashonly)
- **qdance**: [Add extractor](https://github.com/yt-dlp/yt-dlp/commit/f0a1ff118145b6449982ba401f9a9f656ecd8062) ([#7420](https://github.com/yt-dlp/yt-dlp/issues/7420)) by [bashonly](https://github.com/bashonly)
- **sbs**: [Python 3.7 compat](https://github.com/yt-dlp/yt-dlp/commit/f393bbe724b1fc6c7f754a5da507e807b2b40ad2) by [pukkandan](https://github.com/pukkandan)
- **stacommu**: [Add extractors](https://github.com/yt-dlp/yt-dlp/commit/af1fd12f675220df6793fc019dff320bc76e8080) ([#7432](https://github.com/yt-dlp/yt-dlp/issues/7432)) by [urectanc](https://github.com/urectanc)
- **twitter**
    - [Fix unauthenticated extraction](https://github.com/yt-dlp/yt-dlp/commit/49296437a8e5fa91dacb5446e51ab588474c85d3) ([#7476](https://github.com/yt-dlp/yt-dlp/issues/7476)) by [bashonly](https://github.com/bashonly)
    - spaces: [Fix extraction](https://github.com/yt-dlp/yt-dlp/commit/1cffd621cb371f1563563cfb2fe37d137e8a7bee) ([#7512](https://github.com/yt-dlp/yt-dlp/issues/7512)) by [bashonly](https://github.com/bashonly)
- **vidlii**: [Handle relative URLs](https://github.com/yt-dlp/yt-dlp/commit/ad8902f616ad2541f9b9626738f1393fad89a64c) by [pukkandan](https://github.com/pukkandan)
- **vk**: VKPlay, VKPlayLive: [Add extractors](https://github.com/yt-dlp/yt-dlp/commit/8776349ef6b1f644584a92dfa00a05208a48edc4) ([#7358](https://github.com/yt-dlp/yt-dlp/issues/7358)) by [c-basalt](https://github.com/c-basalt)
- **youtube**
    - [Add extractor-arg `formats`](https://github.com/yt-dlp/yt-dlp/commit/58786a10f212bd63f9ad1d0b4d9e4d31c3b385e2) by [pukkandan](https://github.com/pukkandan)
    - [Avoid false DRM detection](https://github.com/yt-dlp/yt-dlp/commit/94ed638a437fc766699d440e978982e24ce6a30a) ([#7396](https://github.com/yt-dlp/yt-dlp/issues/7396)) by [pukkandan](https://github.com/pukkandan)
    - [Fix comments' `is_favorited`](https://github.com/yt-dlp/yt-dlp/commit/89bed013741a776506f60380b7fd89d27d0710b4) ([#7390](https://github.com/yt-dlp/yt-dlp/issues/7390)) by [bbilly1](https://github.com/bbilly1)
    - [Ignore incomplete data for comment threads by default](https://github.com/yt-dlp/yt-dlp/commit/4dc4d8473c085900edc841c87c20041233d25b1f) ([#7475](https://github.com/yt-dlp/yt-dlp/issues/7475)) by [coletdjnz](https://github.com/coletdjnz)
    - [Process `post_live` over 2 hours](https://github.com/yt-dlp/yt-dlp/commit/d949c10c45bfc359bdacd52e6a180169b8128958) by [pukkandan](https://github.com/pukkandan)
    - stories: [Remove](https://github.com/yt-dlp/yt-dlp/commit/90db9a3c00ca80492c6a58c542e4cbf4c2710866) ([#7459](https://github.com/yt-dlp/yt-dlp/issues/7459)) by [pukkandan](https://github.com/pukkandan)
    - tab: [Support shorts-only playlists](https://github.com/yt-dlp/yt-dlp/commit/fcbc9ed760be6e3455bbadfaf277b4504b06f068) ([#7425](https://github.com/yt-dlp/yt-dlp/issues/7425)) by [coletdjnz](https://github.com/coletdjnz)

#### Downloader changes
- **aria2c**: [Add `--no-conf`](https://github.com/yt-dlp/yt-dlp/commit/8a8af356e3bba98a7f7d333aff0777d5d92130c8) by [pukkandan](https://github.com/pukkandan)
- **external**: [Scope cookies](https://github.com/yt-dlp/yt-dlp/commit/1ceb657bdd254ad961489e5060f2ccc7d556b729) by [bashonly](https://github.com/bashonly), [coletdjnz](https://github.com/coletdjnz)
- **http**: [Avoid infinite loop when no data is received](https://github.com/yt-dlp/yt-dlp/commit/662ef1e910b72e57957f06589925b2332ba52821) by [pukkandan](https://github.com/pukkandan)

#### Misc. changes
- [Add CodeQL workflow](https://github.com/yt-dlp/yt-dlp/commit/6355b5f1e1e8e7f4ef866d71d51e03baf0e82f17) ([#7497](https://github.com/yt-dlp/yt-dlp/issues/7497)) by [jorgectf](https://github.com/jorgectf)
- **cleanup**: Miscellaneous: [337734d](https://github.com/yt-dlp/yt-dlp/commit/337734d4a8a6500bc65434843db346b5cbd05e81) by [pukkandan](https://github.com/pukkandan)
- **docs**: [Minor fixes](https://github.com/yt-dlp/yt-dlp/commit/b532a3481046e1eabb6232ee8196fb696c356ff6) by [pukkandan](https://github.com/pukkandan)
- **make_changelog**: [Skip reverted commits](https://github.com/yt-dlp/yt-dlp/commit/fa44802809d189fca0f4782263d48d6533384503) by [pukkandan](https://github.com/pukkandan)

### 2023.06.22

#### Core changes
- [Fix bug in db3ad8a67661d7b234a6954d9c6a4a9b1749f5eb](https://github.com/yt-dlp/yt-dlp/commit/d7cd97e8d8d42b500fea9abb2aa4ac9b0f98b2ad) by [pukkandan](https://github.com/pukkandan)
- [Improve `--download-sections`](https://github.com/yt-dlp/yt-dlp/commit/b4e0d75848e9447cee2cd3646ce54d4744a7ff56) by [pukkandan](https://github.com/pukkandan)
    - Support negative time-ranges
    - Add `*from-url` to obey time-ranges in URL
- [Indicate `filesize` approximated from `tbr` better](https://github.com/yt-dlp/yt-dlp/commit/0dff8e4d1e6e9fb938f4256ea9af7d81f42fd54f) by [pukkandan](https://github.com/pukkandan)

#### Extractor changes
- [Support multiple `_VALID_URL`s](https://github.com/yt-dlp/yt-dlp/commit/5fd8367496b42c7b900b896a0d5460561a2859de) ([#5812](https://github.com/yt-dlp/yt-dlp/issues/5812)) by [nixxo](https://github.com/nixxo)
- **dplay**: GlobalCyclingNetworkPlus: [Add extractor](https://github.com/yt-dlp/yt-dlp/commit/774aa09dd6aa61ced9ec818d1f67e53414d22762) ([#7360](https://github.com/yt-dlp/yt-dlp/issues/7360)) by [bashonly](https://github.com/bashonly)
- **dropout**: [Fix season extraction](https://github.com/yt-dlp/yt-dlp/commit/db22142f6f817ff673d417b4b78e8db497bf8ab3) ([#7304](https://github.com/yt-dlp/yt-dlp/issues/7304)) by [OverlordQ](https://github.com/OverlordQ)
- **motherless**: [Add gallery support, fix groups](https://github.com/yt-dlp/yt-dlp/commit/f2ff0f6f1914b82d4a51681a72cc0828115dcb4a) ([#7211](https://github.com/yt-dlp/yt-dlp/issues/7211)) by [rexlambert22](https://github.com/rexlambert22), [Ti4eeT4e](https://github.com/Ti4eeT4e)
- **nebula**: [Fix extractor](https://github.com/yt-dlp/yt-dlp/commit/3f756c8c4095b942cf49788eb0862ceaf57847f2) ([#7156](https://github.com/yt-dlp/yt-dlp/issues/7156)) by [Lamieur](https://github.com/Lamieur), [rohieb](https://github.com/rohieb)
- **rheinmaintv**: [Add extractor](https://github.com/yt-dlp/yt-dlp/commit/98cb1eda7a4cf67c96078980dbd63e6c06ad7f7c) ([#7311](https://github.com/yt-dlp/yt-dlp/issues/7311)) by [barthelmannk](https://github.com/barthelmannk)
- **youtube**
    - [Add `ios` to default clients used](https://github.com/yt-dlp/yt-dlp/commit/1e75d97db21152acc764b30a688e516f04b8a142) by [pukkandan](https://github.com/pukkandan)
        - IOS is affected neither by 403 nor by nsig so helps mitigate them preemptively
        - IOS also has higher bit-rate 'premium' formats though they are not labeled as such
    - [Improve description parsing performance](https://github.com/yt-dlp/yt-dlp/commit/71dc18fa29263a1ff0472c23d81bfc8dd4422d48) ([#7315](https://github.com/yt-dlp/yt-dlp/issues/7315)) by [berkanteber](https://github.com/berkanteber), [pukkandan](https://github.com/pukkandan)
    - [Improve nsig function name extraction](https://github.com/yt-dlp/yt-dlp/commit/cd810afe2ac5567c822b7424800fc470ef2d0045) by [pukkandan](https://github.com/pukkandan)
    - [Workaround 403 for android formats](https://github.com/yt-dlp/yt-dlp/commit/81ca451480051d7ce1a31c017e005358345a9149) by [pukkandan](https://github.com/pukkandan)

#### Misc. changes
- [Revert "Add automatic duplicate issue detection"](https://github.com/yt-dlp/yt-dlp/commit/a4486bfc1dc7057efca9dd3fe70d7fa25c56f700) by [pukkandan](https://github.com/pukkandan)
- **cleanup**
    - Miscellaneous
        - [7f9c6a6](https://github.com/yt-dlp/yt-dlp/commit/7f9c6a63b16e145495479e9f666f5b9e2ee69e2f) by [bashonly](https://github.com/bashonly)
        - [812cdfa](https://github.com/yt-dlp/yt-dlp/commit/812cdfa06c33a40e73a8e04b3e6f42c084666a43) by [pukkandan](https://github.com/pukkandan)

### 2023.06.21

#### Important changes
- YouTube: Improved throttling and signature fixes

#### Core changes
- [Add `--compat-option playlist-match-filter`](https://github.com/yt-dlp/yt-dlp/commit/93b39cdbd9dcf351bfa0c4ee252805b4617fdca9) by [pukkandan](https://github.com/pukkandan)
- [Add `--no-quiet`](https://github.com/yt-dlp/yt-dlp/commit/d669772c65e8630162fd6555d0a578b246591921) by [pukkandan](https://github.com/pukkandan)
- [Add option `--color`](https://github.com/yt-dlp/yt-dlp/commit/8417f26b8a819cd7ffcd4e000ca3e45033e670fb) ([#6904](https://github.com/yt-dlp/yt-dlp/issues/6904)) by [Grub4K](https://github.com/Grub4K)
- [Add option `--netrc-cmd`](https://github.com/yt-dlp/yt-dlp/commit/db3ad8a67661d7b234a6954d9c6a4a9b1749f5eb) ([#6682](https://github.com/yt-dlp/yt-dlp/issues/6682)) by [NDagestad](https://github.com/NDagestad), [pukkandan](https://github.com/pukkandan)
- [Add option `--xff`](https://github.com/yt-dlp/yt-dlp/commit/c16644642b08e2bf4130a6c5fa01395d8718c990) by [pukkandan](https://github.com/pukkandan)
- [Auto-select default format in `-f-`](https://github.com/yt-dlp/yt-dlp/commit/372a0f3b9dadd1e52234b498aa4c7040ef868c7d) ([#7101](https://github.com/yt-dlp/yt-dlp/issues/7101)) by [ivanskodje](https://github.com/ivanskodje), [pukkandan](https://github.com/pukkandan)
- [Deprecate internal `Youtubedl-no-compression` header](https://github.com/yt-dlp/yt-dlp/commit/955c89584b66fcd0fcfab3e611f1edeb1ca63886) ([#6876](https://github.com/yt-dlp/yt-dlp/issues/6876)) by [coletdjnz](https://github.com/coletdjnz)
- [Do not translate newlines in `--print-to-file`](https://github.com/yt-dlp/yt-dlp/commit/9874e82b5a61582169300bea561b3e8899ad1ef7) by [pukkandan](https://github.com/pukkandan)
- [Ensure pre-processor errors do not block `--print`](https://github.com/yt-dlp/yt-dlp/commit/f005a35aa7e4f67a0c603a946c0dd714c151b2d6) by [pukkandan](https://github.com/pukkandan) (With fixes in [17ba434](https://github.com/yt-dlp/yt-dlp/commit/17ba4343cf99701692a7f4798fd42b50f644faba))
- [Fix `filepath` being copied to underlying format dict](https://github.com/yt-dlp/yt-dlp/commit/84078a8b38f403495d00b46654c8750774d821de) by [pukkandan](https://github.com/pukkandan)
- [Improve HTTP redirect handling](https://github.com/yt-dlp/yt-dlp/commit/08916a49c777cb6e000eec092881eb93ec22076c) ([#7094](https://github.com/yt-dlp/yt-dlp/issues/7094)) by [coletdjnz](https://github.com/coletdjnz)
- [Populate `filename` and `urls` fields at all stages of `--print`](https://github.com/yt-dlp/yt-dlp/commit/170605840ea9d5ad75da6576485ea7d125b428ee) by [pukkandan](https://github.com/pukkandan) (With fixes in [b5f61b6](https://github.com/yt-dlp/yt-dlp/commit/b5f61b69d4561b81fc98c226b176f0c15493e688))
- [Relaxed validation for numeric format filters](https://github.com/yt-dlp/yt-dlp/commit/c3f624ef0a5d7a6ae1c5ffeb243087e9fc7d79dc) by [pukkandan](https://github.com/pukkandan)
- [Support decoding multiple content encodings](https://github.com/yt-dlp/yt-dlp/commit/daafbf49b3482edae4d70dd37070be99742a926e) ([#7142](https://github.com/yt-dlp/yt-dlp/issues/7142)) by [coletdjnz](https://github.com/coletdjnz)
- [Support loading info.json with a list at it's root](https://github.com/yt-dlp/yt-dlp/commit/ab1de9cb1e39cf421c2b7dc6756c6ff1955bb313) by [pukkandan](https://github.com/pukkandan)
- [Workaround erroneous urllib Windows proxy parsing](https://github.com/yt-dlp/yt-dlp/commit/3f66b6fe50f8d5b545712f8b19d5ae62f5373980) ([#7092](https://github.com/yt-dlp/yt-dlp/issues/7092)) by [coletdjnz](https://github.com/coletdjnz)
- **cookies**
    - [Defer extraction of v11 key from keyring](https://github.com/yt-dlp/yt-dlp/commit/9b7a48abd1b187eae1e3f6c9839c47d43ccec00b) by [Grub4K](https://github.com/Grub4K)
    - [Move `YoutubeDLCookieJar` to cookies module](https://github.com/yt-dlp/yt-dlp/commit/b87e01c123fd560b6a674ce00f45a9459d82d98a) ([#7091](https://github.com/yt-dlp/yt-dlp/issues/7091)) by [coletdjnz](https://github.com/coletdjnz)
    - [Support custom Safari cookies path](https://github.com/yt-dlp/yt-dlp/commit/a58182b75a05fe0a10c5e94a536711d3ade19c20) ([#6783](https://github.com/yt-dlp/yt-dlp/issues/6783)) by [NextFire](https://github.com/NextFire)
    - [Update for chromium changes](https://github.com/yt-dlp/yt-dlp/commit/b38d4c941d1993ab27e4c0f8e024e23c2ec0f8f8) ([#6897](https://github.com/yt-dlp/yt-dlp/issues/6897)) by [mbway](https://github.com/mbway)
- **Cryptodome**: [Fix `__bool__`](https://github.com/yt-dlp/yt-dlp/commit/98ac902c4979e4529b166e873473bef42baa2e3e) by [pukkandan](https://github.com/pukkandan)
- **jsinterp**
    - [Do not compile regex](https://github.com/yt-dlp/yt-dlp/commit/7aeda6cc9e73ada0b0a0b6a6748c66bef63a20a8) by [pukkandan](https://github.com/pukkandan)
    - [Fix division](https://github.com/yt-dlp/yt-dlp/commit/b4a252fba81f53631c07ca40ce7583f5d19a8a36) ([#7279](https://github.com/yt-dlp/yt-dlp/issues/7279)) by [bashonly](https://github.com/bashonly)
    - [Fix global object extraction](https://github.com/yt-dlp/yt-dlp/commit/01aba2519a0884ef17d5f85608dbd2a455577147) by [pukkandan](https://github.com/pukkandan)
    - [Handle `NaN` in bitwise operators](https://github.com/yt-dlp/yt-dlp/commit/1d7656184c6b8aa46b29149893894b3c24f1df00) by [pukkandan](https://github.com/pukkandan)
    - [Handle negative numbers better](https://github.com/yt-dlp/yt-dlp/commit/7cf51f21916292cd80bdeceb37489f5322f166dd) by [pukkandan](https://github.com/pukkandan)
- **outtmpl**
    - [Allow `\n` in replacements and default.](https://github.com/yt-dlp/yt-dlp/commit/78fde6e3398ff11e5d383a66b28664badeab5180) by [pukkandan](https://github.com/pukkandan)
    - [Fix some minor bugs](https://github.com/yt-dlp/yt-dlp/commit/ebe1b4e34f43c3acad30e4bcb8484681a030c114) by [pukkandan](https://github.com/pukkandan) (With fixes in [1619ab3](https://github.com/yt-dlp/yt-dlp/commit/1619ab3e67d8dc4f86fc7ed292c79345bc0d91a0))
    - [Support `str.format` syntax inside replacements](https://github.com/yt-dlp/yt-dlp/commit/ec9311c41b111110bc52cfbd6ea682c6fb23f77a) by [pukkandan](https://github.com/pukkandan)
- **update**
    - [Better error handling](https://github.com/yt-dlp/yt-dlp/commit/d2e84d5eb01c66fc5304e8566348d65a7be24ed7) by [pukkandan](https://github.com/pukkandan)
    - [Do not restart into versions without `--update-to`](https://github.com/yt-dlp/yt-dlp/commit/02948a17d903f544363bb20b51a6d8baed7bba08) by [pukkandan](https://github.com/pukkandan)
    - [Implement `--update-to` repo](https://github.com/yt-dlp/yt-dlp/commit/665472a7de3880578c0b7b3f95c71570c056368e) by [Grub4K](https://github.com/Grub4K), [pukkandan](https://github.com/pukkandan)
- **upstream**
    - [Merged with youtube-dl 07af47](https://github.com/yt-dlp/yt-dlp/commit/42f2d40b475db66486a4b4fe5b56751a640db5db) by [pukkandan](https://github.com/pukkandan)
    - [Merged with youtube-dl d1c6c5](https://github.com/yt-dlp/yt-dlp/commit/4823ec9f461512daa1b8ab362893bb86a6320b26) by [pukkandan](https://github.com/pukkandan) (With fixes in [edbe5b5](https://github.com/yt-dlp/yt-dlp/commit/edbe5b589dd0860a67b4e03f58db3cd2539d91c2) by [bashonly](https://github.com/bashonly))
- **utils**
    - `FormatSorter`: [Improve `size` and `br`](https://github.com/yt-dlp/yt-dlp/commit/eedda5252c05327748dede204a8fccafa0288118) by [pukkandan](https://github.com/pukkandan), [u-spec-png](https://github.com/u-spec-png)
    - `js_to_json`: [Implement template strings](https://github.com/yt-dlp/yt-dlp/commit/0898c5c8ccadfc404472456a7a7751b72afebadd) ([#6623](https://github.com/yt-dlp/yt-dlp/issues/6623)) by [Grub4K](https://github.com/Grub4K)
    - `locked_file`: [Fix for virtiofs](https://github.com/yt-dlp/yt-dlp/commit/45998b3e371b819ce0dbe50da703809a048cc2fe) ([#6840](https://github.com/yt-dlp/yt-dlp/issues/6840)) by [brandon-dacrib](https://github.com/brandon-dacrib)
    - `strftime_or_none`: [Handle negative timestamps](https://github.com/yt-dlp/yt-dlp/commit/a35af4306d24c56c6358f89cdf204860d1cd62b4) by [dirkf](https://github.com/dirkf), [pukkandan](https://github.com/pukkandan)
    - `traverse_obj`
        - [Allow iterables in traversal](https://github.com/yt-dlp/yt-dlp/commit/21b5ec86c2c37d10c5bb97edd7051d3aac16bb3e) ([#6902](https://github.com/yt-dlp/yt-dlp/issues/6902)) by [Grub4K](https://github.com/Grub4K)
        - [More fixes](https://github.com/yt-dlp/yt-dlp/commit/b079c26f0af8085bccdadc72c61c8164ca5ab0f8) ([#6959](https://github.com/yt-dlp/yt-dlp/issues/6959)) by [Grub4K](https://github.com/Grub4K)
    - `write_string`: [Fix noconsole behavior](https://github.com/yt-dlp/yt-dlp/commit/3b479100df02e20dd949e046003ae96ddbfced57) by [Grub4K](https://github.com/Grub4K)

#### Extractor changes
- [Do not exit early for unsuitable `url_result`](https://github.com/yt-dlp/yt-dlp/commit/baa922b5c74b10e3b86ff5e6cf6529b3aae8efab) by [pukkandan](https://github.com/pukkandan)
- [Do not warn for invalid chapter data in description](https://github.com/yt-dlp/yt-dlp/commit/84ffeb7d5e72e3829319ba7720a8480fc4c7503b) by [pukkandan](https://github.com/pukkandan)
- [Extract more metadata from ISM](https://github.com/yt-dlp/yt-dlp/commit/f68434cc74cfd3db01b266476a2eac8329fbb267) by [pukkandan](https://github.com/pukkandan)
- **abematv**: [Add fallback for title and description extraction and extract more metadata](https://github.com/yt-dlp/yt-dlp/commit/c449c0655d7c8549e6e1389c26b628053b253d39) ([#6994](https://github.com/yt-dlp/yt-dlp/issues/6994)) by [Lesmiscore](https://github.com/Lesmiscore)
- **acast**: [Support embeds](https://github.com/yt-dlp/yt-dlp/commit/c91ac833ea99b00506e470a44cf930e4e23378c9) ([#7212](https://github.com/yt-dlp/yt-dlp/issues/7212)) by [pabs3](https://github.com/pabs3)
- **adobepass**: [Handle `Charter_Direct` MSO as `Spectrum`](https://github.com/yt-dlp/yt-dlp/commit/ea0570820336a0fe9c3b530d1b0d1e59313274f4) ([#6824](https://github.com/yt-dlp/yt-dlp/issues/6824)) by [bashonly](https://github.com/bashonly)
- **aeonco**: [Support Youtube embeds](https://github.com/yt-dlp/yt-dlp/commit/ed81b74802b4247ee8d9dc0ef87eb52baefede1c) ([#6591](https://github.com/yt-dlp/yt-dlp/issues/6591)) by [alexklapheke](https://github.com/alexklapheke)
- **afreecatv**: [Fix extractor](https://github.com/yt-dlp/yt-dlp/commit/fdd69db38924c38194ef236b26325d66ac815c88) ([#6283](https://github.com/yt-dlp/yt-dlp/issues/6283)) by [blmarket](https://github.com/blmarket)
- **ARDBetaMediathek**: [Add thumbnail](https://github.com/yt-dlp/yt-dlp/commit/f78eb41e1c0f1dcdb10317358a26bf541dc7ee15) ([#6890](https://github.com/yt-dlp/yt-dlp/issues/6890)) by [StefanLobbenmeier](https://github.com/StefanLobbenmeier)
- **bibeltv**: [Fix extraction, support live streams and series](https://github.com/yt-dlp/yt-dlp/commit/4ad58667c102bd82a7c4cca8aa395ec1682e3b4c) ([#6505](https://github.com/yt-dlp/yt-dlp/issues/6505)) by [flashdagger](https://github.com/flashdagger)
- **bilibili**
    - [Support festival videos](https://github.com/yt-dlp/yt-dlp/commit/ab29e47029e2f5b48abbbab78e82faf7cf6e9506) ([#6547](https://github.com/yt-dlp/yt-dlp/issues/6547)) by [qbnu](https://github.com/qbnu)
    - SpaceVideo: [Extract signature](https://github.com/yt-dlp/yt-dlp/commit/6f10cdcf7eeaeae5b75e0a4428cd649c156a2d83) ([#7149](https://github.com/yt-dlp/yt-dlp/issues/7149)) by [elyse0](https://github.com/elyse0)
- **biliIntl**: [Add comment extraction](https://github.com/yt-dlp/yt-dlp/commit/b093c38cc9f26b59a8504211d792f053142c847d) ([#6079](https://github.com/yt-dlp/yt-dlp/issues/6079)) by [HobbyistDev](https://github.com/HobbyistDev)
- **bitchute**: [Add more fallback subdomains](https://github.com/yt-dlp/yt-dlp/commit/0c4e0fbcade0fc92d14c2a6d63e360fe067f6192) ([#6907](https://github.com/yt-dlp/yt-dlp/issues/6907)) by [Neurognostic](https://github.com/Neurognostic)
- **booyah**: [Remove extractor](https://github.com/yt-dlp/yt-dlp/commit/f7f7a877bf8e87fd4eb0ad2494ad948ca7691114) by [pukkandan](https://github.com/pukkandan)
- **BrainPOP**: [Add extractors](https://github.com/yt-dlp/yt-dlp/commit/979568f26ece80bca72b48f0dd57d676e431059a) ([#6106](https://github.com/yt-dlp/yt-dlp/issues/6106)) by [MinePlayersPE](https://github.com/MinePlayersPE)
- **bravotv**
    - [Detect DRM](https://github.com/yt-dlp/yt-dlp/commit/1fe5bf240e6ade487d18079a62aa36bcc440a27a) ([#7171](https://github.com/yt-dlp/yt-dlp/issues/7171)) by [bashonly](https://github.com/bashonly)
    - [Fix extractor](https://github.com/yt-dlp/yt-dlp/commit/06966cb8966b9aa4f60ab9c44c182a057d4ca3a3) ([#6568](https://github.com/yt-dlp/yt-dlp/issues/6568)) by [bashonly](https://github.com/bashonly)
- **camfm**: [Add extractors](https://github.com/yt-dlp/yt-dlp/commit/4cbfa570a1b9bd65b0f48770693377e8d842dcb0) ([#7083](https://github.com/yt-dlp/yt-dlp/issues/7083)) by [garret1317](https://github.com/garret1317)
- **cbc**
    - [Fix live extractor, playlist `_VALID_URL`](https://github.com/yt-dlp/yt-dlp/commit/7a7b1376fbce0067cf37566bb47131bc0022638d) ([#6625](https://github.com/yt-dlp/yt-dlp/issues/6625)) by [makew0rld](https://github.com/makew0rld)
    - [Ignore 426 from API](https://github.com/yt-dlp/yt-dlp/commit/4afb208cf07b59291ae3b0c4efc83945ee5b8812) ([#6781](https://github.com/yt-dlp/yt-dlp/issues/6781)) by [jo-nike](https://github.com/jo-nike)
    - gem: [Update `_VALID_URL`](https://github.com/yt-dlp/yt-dlp/commit/871c907454693940cb56906ed9ea49fcb7154829) ([#6499](https://github.com/yt-dlp/yt-dlp/issues/6499)) by [makeworld-the-better-one](https://github.com/makeworld-the-better-one)
- **cbs**: [Add `ParamountPressExpress` extractor](https://github.com/yt-dlp/yt-dlp/commit/44369c9afa996e14e9f466754481d878811b5b4a) ([#6604](https://github.com/yt-dlp/yt-dlp/issues/6604)) by [bashonly](https://github.com/bashonly)
- **cbsnews**: [Overhaul extractors](https://github.com/yt-dlp/yt-dlp/commit/f6e43d6fa9804c24525e1fed0a87782754dab7ed) ([#6681](https://github.com/yt-dlp/yt-dlp/issues/6681)) by [bashonly](https://github.com/bashonly)
- **chilloutzone**: [Fix extractor](https://github.com/yt-dlp/yt-dlp/commit/6f4fc5660f40f3458882a8f51601eae4af7be609) ([#6445](https://github.com/yt-dlp/yt-dlp/issues/6445)) by [bashonly](https://github.com/bashonly)
- **clipchamp**: [Add extractor](https://github.com/yt-dlp/yt-dlp/commit/2f07c4c1da4361af213e5791279b9d152d2e4ce3) ([#6978](https://github.com/yt-dlp/yt-dlp/issues/6978)) by [bashonly](https://github.com/bashonly)
- **comedycentral**: [Add support for movies](https://github.com/yt-dlp/yt-dlp/commit/66468bbf49562ff82670cbbd456c5e8448a6df34) ([#7108](https://github.com/yt-dlp/yt-dlp/issues/7108)) by [sqrtNOT](https://github.com/sqrtNOT)
- **crtvg**: [Add extractor](https://github.com/yt-dlp/yt-dlp/commit/26c517b29c8727e47948d6fff749d5297f0efb60) ([#7168](https://github.com/yt-dlp/yt-dlp/issues/7168)) by [ItzMaxTV](https://github.com/ItzMaxTV)
- **crunchyroll**: [Rework with support for movies, music and artists](https://github.com/yt-dlp/yt-dlp/commit/032de83ea9ff2f4977d9c71a93bbc1775597b762) ([#6237](https://github.com/yt-dlp/yt-dlp/issues/6237)) by [Grub4K](https://github.com/Grub4K)
- **dacast**: [Add extractors](https://github.com/yt-dlp/yt-dlp/commit/c25cac2f8e5fbac2737a426d7778fd2f0efc5381) ([#6896](https://github.com/yt-dlp/yt-dlp/issues/6896)) by [bashonly](https://github.com/bashonly)
- **daftsex**: [Update domain and embed player url](https://github.com/yt-dlp/yt-dlp/commit/fc5a7f9b27d2a89b1f3ca7d33a95301c21d832cd) ([#5966](https://github.com/yt-dlp/yt-dlp/issues/5966)) by [JChris246](https://github.com/JChris246)
- **DigitalConcertHall**: [Support films](https://github.com/yt-dlp/yt-dlp/commit/55ed4ff73487feb3177b037dfc2ea527e777da3e) ([#7202](https://github.com/yt-dlp/yt-dlp/issues/7202)) by [ItzMaxTV](https://github.com/ItzMaxTV)
- **discogs**: [Add extractor](https://github.com/yt-dlp/yt-dlp/commit/6daaf21092888beff11b807cd46f832f1f9c46a0) ([#6624](https://github.com/yt-dlp/yt-dlp/issues/6624)) by [rjy](https://github.com/rjy)
- **dlf**: [Add extractors](https://github.com/yt-dlp/yt-dlp/commit/b423b6a48e0b19260bc95ab7d72d2138d7f124dc) ([#6697](https://github.com/yt-dlp/yt-dlp/issues/6697)) by [nick-cd](https://github.com/nick-cd)
- **drtv**: [Fix radio page extraction](https://github.com/yt-dlp/yt-dlp/commit/9a06b7b1891b48cebbe275652ae8025a36d97d97) ([#6552](https://github.com/yt-dlp/yt-dlp/issues/6552)) by [viktor-enzell](https://github.com/viktor-enzell)
- **Dumpert**: [Fix m3u8 and support new URL pattern](https://github.com/yt-dlp/yt-dlp/commit/f8ae441501596733e2b967430471643a1d7cacb8) ([#6091](https://github.com/yt-dlp/yt-dlp/issues/6091)) by [DataGhost](https://github.com/DataGhost), [pukkandan](https://github.com/pukkandan)
- **elevensports**: [Add extractor](https://github.com/yt-dlp/yt-dlp/commit/ecfe47973f6603b5367fe2cc3c65274627d94516) ([#7172](https://github.com/yt-dlp/yt-dlp/issues/7172)) by [ItzMaxTV](https://github.com/ItzMaxTV)
- **ettutv**: [Add extractor](https://github.com/yt-dlp/yt-dlp/commit/83465fc4100a2fb2c188898fbc2f3021f6a9b4dd) ([#6579](https://github.com/yt-dlp/yt-dlp/issues/6579)) by [elyse0](https://github.com/elyse0)
- **europarl**: [Rewrite extractor](https://github.com/yt-dlp/yt-dlp/commit/03789976d301eaed3e957dbc041573098f6af059) ([#7114](https://github.com/yt-dlp/yt-dlp/issues/7114)) by [HobbyistDev](https://github.com/HobbyistDev)
- **eurosport**: [Improve `_VALID_URL`](https://github.com/yt-dlp/yt-dlp/commit/45e87ea106ad37b2a002663fa30ee41ce97b16cd) ([#7076](https://github.com/yt-dlp/yt-dlp/issues/7076)) by [HobbyistDev](https://github.com/HobbyistDev)
- **facebook**: [Fix metadata extraction](https://github.com/yt-dlp/yt-dlp/commit/3b52a606881e6adadc33444abdeacce562b79330) ([#6856](https://github.com/yt-dlp/yt-dlp/issues/6856)) by [ringus1](https://github.com/ringus1)
- **foxnews**: [Fix extractors](https://github.com/yt-dlp/yt-dlp/commit/97d60ad8cd6c99f01e463a9acfce8693aff2a609) ([#7222](https://github.com/yt-dlp/yt-dlp/issues/7222)) by [bashonly](https://github.com/bashonly)
- **funker530**: [Add extractor](https://github.com/yt-dlp/yt-dlp/commit/cab94a0cd8b6d3fffed5a6faff030274adbed182) ([#7291](https://github.com/yt-dlp/yt-dlp/issues/7291)) by [Cyberes](https://github.com/Cyberes)
- **generic**
    - [Accept values for `fragment_query`, `variant_query`](https://github.com/yt-dlp/yt-dlp/commit/5cc0a8fd2e9fec50026fb92170b57993af939e4a) ([#6600](https://github.com/yt-dlp/yt-dlp/issues/6600)) by [bashonly](https://github.com/bashonly) (With fixes in [9bfe0d1](https://github.com/yt-dlp/yt-dlp/commit/9bfe0d15bd7dbdc6b0e6378fa9f5e2e289b2373b))
    - [Add extractor-args `hls_key`, `variant_query`](https://github.com/yt-dlp/yt-dlp/commit/c2e0fc40a73dd85ab3920f977f579d475e66ef59) ([#6567](https://github.com/yt-dlp/yt-dlp/issues/6567)) by [bashonly](https://github.com/bashonly)
    - [Attempt to detect live HLS](https://github.com/yt-dlp/yt-dlp/commit/93e7c6995e07dafb9dcc06c0d06acf6c5bdfecc5) ([#6775](https://github.com/yt-dlp/yt-dlp/issues/6775)) by [bashonly](https://github.com/bashonly)
- **genius**: [Add support for articles](https://github.com/yt-dlp/yt-dlp/commit/460da07439718d9af1e3661da2a23e05a913a2e6) ([#6474](https://github.com/yt-dlp/yt-dlp/issues/6474)) by [bashonly](https://github.com/bashonly)
- **globalplayer**: [Add extractors](https://github.com/yt-dlp/yt-dlp/commit/30647668a92a0ca5cd108776804baac0996bd9f7) ([#6903](https://github.com/yt-dlp/yt-dlp/issues/6903)) by [garret1317](https://github.com/garret1317)
- **gmanetwork**: [Add extractor](https://github.com/yt-dlp/yt-dlp/commit/2d97d154fe4fb84fe2ed3a4e1ed5819e89b71e88) ([#5945](https://github.com/yt-dlp/yt-dlp/issues/5945)) by [HobbyistDev](https://github.com/HobbyistDev)
- **gronkh**: [Extract duration and chapters](https://github.com/yt-dlp/yt-dlp/commit/9c92b803fa24e48543ce969468d5404376e315b7) ([#6817](https://github.com/yt-dlp/yt-dlp/issues/6817)) by [satan1st](https://github.com/satan1st)
- **hentaistigma**: [Remove extractor](https://github.com/yt-dlp/yt-dlp/commit/04f8018a0544736a18494bc3899d06b05b78fae6) by [pukkandan](https://github.com/pukkandan)
- **hidive**: [Fix login](https://github.com/yt-dlp/yt-dlp/commit/e6ab678e36c40ded0aae305bbb866cdab554d417) by [pukkandan](https://github.com/pukkandan)
- **hollywoodreporter**: [Add extractors](https://github.com/yt-dlp/yt-dlp/commit/6bdb64e2a2a6d504d8ce1dc830fbfb8a7f199c63) ([#6614](https://github.com/yt-dlp/yt-dlp/issues/6614)) by [bashonly](https://github.com/bashonly)
- **hotstar**: [Support `/shows/` URLs](https://github.com/yt-dlp/yt-dlp/commit/7f8ddebbb51c9fd4a347306332a718ba41b371b8) ([#7225](https://github.com/yt-dlp/yt-dlp/issues/7225)) by [bashonly](https://github.com/bashonly)
- **hrefli**: [Add extractor](https://github.com/yt-dlp/yt-dlp/commit/7e35526d5b970a034b9d76215ee3e4bd7631edcd) ([#6762](https://github.com/yt-dlp/yt-dlp/issues/6762)) by [selfisekai](https://github.com/selfisekai)
- **idolplus**: [Add extractor](https://github.com/yt-dlp/yt-dlp/commit/5c14b213679ed4401288bdc86ae696932e219222) ([#6732](https://github.com/yt-dlp/yt-dlp/issues/6732)) by [ping](https://github.com/ping)
- **iq**: [Set more language codes](https://github.com/yt-dlp/yt-dlp/commit/2d5cae9636714ff922d28c548c349d5f2b48f317) ([#6476](https://github.com/yt-dlp/yt-dlp/issues/6476)) by [D0LLYNH0](https://github.com/D0LLYNH0)
- **iwara**
    - [Accept old URLs](https://github.com/yt-dlp/yt-dlp/commit/ab92d8651c48d247dfb7d3f0a824cc986e47c7ed) by [Lesmiscore](https://github.com/Lesmiscore)
    - [Fix authentication](https://github.com/yt-dlp/yt-dlp/commit/0a5d7c39e17bb9bd50c9db42bcad40eb82d7f784) ([#7137](https://github.com/yt-dlp/yt-dlp/issues/7137)) by [toomyzoom](https://github.com/toomyzoom)
    - [Fix format sorting](https://github.com/yt-dlp/yt-dlp/commit/56793f74c36899742d7abd52afb0deca97d469e1) ([#6651](https://github.com/yt-dlp/yt-dlp/issues/6651)) by [hasezoey](https://github.com/hasezoey)
    - [Fix typo](https://github.com/yt-dlp/yt-dlp/commit/d1483ec693c79f0b4ddf493870bcb840aca4da08) by [Lesmiscore](https://github.com/Lesmiscore)
    - [Implement login](https://github.com/yt-dlp/yt-dlp/commit/21b9413cf7dd4830b2ece57af21589dd4538fc52) ([#6721](https://github.com/yt-dlp/yt-dlp/issues/6721)) by [toomyzoom](https://github.com/toomyzoom)
    - [Overhaul extractors](https://github.com/yt-dlp/yt-dlp/commit/c14af7a741931b364bab3d9546c0f4359f318f8c) ([#6557](https://github.com/yt-dlp/yt-dlp/issues/6557)) by [Lesmiscore](https://github.com/Lesmiscore)
    - [Report private videos](https://github.com/yt-dlp/yt-dlp/commit/95a383be1b6fb00c92ee3fb091732c4f6009acb6) ([#6641](https://github.com/yt-dlp/yt-dlp/issues/6641)) by [Lesmiscore](https://github.com/Lesmiscore)
- **JStream**: [Add extractor](https://github.com/yt-dlp/yt-dlp/commit/3459d3c5af3b2572ed51e8ecfda6c11022a838c6) ([#6252](https://github.com/yt-dlp/yt-dlp/issues/6252)) by [Lesmiscore](https://github.com/Lesmiscore)
- **jwplatform**: [Update `_extract_embed_urls`](https://github.com/yt-dlp/yt-dlp/commit/cf9fd52fabe71d6e7c30d3ea525029ffa561fc9c) ([#6383](https://github.com/yt-dlp/yt-dlp/issues/6383)) by [carusocr](https://github.com/carusocr)
- **kick**: [Make initial request non-fatal](https://github.com/yt-dlp/yt-dlp/commit/0a6918a4a1431960181d8c50e0bbbcb0afbaff9a) by [bashonly](https://github.com/bashonly)
- **LastFM**: [Rewrite playlist extraction](https://github.com/yt-dlp/yt-dlp/commit/026435714cb7c39613a0d7d2acd15d3823b78d94) ([#6379](https://github.com/yt-dlp/yt-dlp/issues/6379)) by [hatienl0i261299](https://github.com/hatienl0i261299), [pukkandan](https://github.com/pukkandan)
- **lbry**: [Extract original quality formats](https://github.com/yt-dlp/yt-dlp/commit/44c0d66442b568d9e1359e669d8b029b08a77fa7) ([#7257](https://github.com/yt-dlp/yt-dlp/issues/7257)) by [bashonly](https://github.com/bashonly)
- **line**: [Remove extractors](https://github.com/yt-dlp/yt-dlp/commit/faa0332ed69e070cf3bd31390589a596e962f392) ([#6734](https://github.com/yt-dlp/yt-dlp/issues/6734)) by [sian1468](https://github.com/sian1468)
- **livestream**: [Support videos with account id](https://github.com/yt-dlp/yt-dlp/commit/bfdf144c7e5d7a93fbfa9d8e65598c72bf2b542a) ([#6324](https://github.com/yt-dlp/yt-dlp/issues/6324)) by [theperfectpunk](https://github.com/theperfectpunk)
- **medaltv**: [Fix clips](https://github.com/yt-dlp/yt-dlp/commit/1e3c2b6ec28d7ab5e31341fa93c47b65be4fbff4) ([#6502](https://github.com/yt-dlp/yt-dlp/issues/6502)) by [xenova](https://github.com/xenova)
- **mediastream**: [Improve `WinSports` and embed extraction](https://github.com/yt-dlp/yt-dlp/commit/03025b6e105139d01cd415ddc51fd692957fd2ba) ([#6426](https://github.com/yt-dlp/yt-dlp/issues/6426)) by [bashonly](https://github.com/bashonly)
- **mgtv**: [Fix formats extraction](https://github.com/yt-dlp/yt-dlp/commit/59d9fe08312bbb76ee26238d207a8ca35410a48d) ([#7234](https://github.com/yt-dlp/yt-dlp/issues/7234)) by [bashonly](https://github.com/bashonly)
- **Mzaalo**: [Add extractor](https://github.com/yt-dlp/yt-dlp/commit/dc3c44f349ba85af320e706e2a27ad81a78b1c6e) ([#7163](https://github.com/yt-dlp/yt-dlp/issues/7163)) by [ItzMaxTV](https://github.com/ItzMaxTV)
- **nbc**: [Fix `NBCStations` direct mp4 formats](https://github.com/yt-dlp/yt-dlp/commit/9be0fe1fd967f62cbf3c60bd14e1021a70abc147) ([#6637](https://github.com/yt-dlp/yt-dlp/issues/6637)) by [bashonly](https://github.com/bashonly)
- **nebula**: [Add `beta.nebula.tv`](https://github.com/yt-dlp/yt-dlp/commit/cbfe2e5cbe0f4649a91e323a82b8f5f774f36662) ([#6516](https://github.com/yt-dlp/yt-dlp/issues/6516)) by [unbeatable-101](https://github.com/unbeatable-101)
- **nekohacker**: [Add extractor](https://github.com/yt-dlp/yt-dlp/commit/489f51279d00318018478fd7461eddbe3b45297e) ([#7003](https://github.com/yt-dlp/yt-dlp/issues/7003)) by [hasezoey](https://github.com/hasezoey)
- **nhk**
    - [Add `NhkRadiru` extractor](https://github.com/yt-dlp/yt-dlp/commit/8f0be90ecb3b8d862397177bb226f17b245ef933) ([#6819](https://github.com/yt-dlp/yt-dlp/issues/6819)) by [garret1317](https://github.com/garret1317)
    - [Fix API extraction](https://github.com/yt-dlp/yt-dlp/commit/f41b949a2ef646fbc36375febbe3f0c19d742c0f) ([#7180](https://github.com/yt-dlp/yt-dlp/issues/7180)) by [menschel](https://github.com/menschel), [sjthespian](https://github.com/sjthespian)
    - `NhkRadiruLive`: [Add extractor](https://github.com/yt-dlp/yt-dlp/commit/81c8b9bdd9841b72cbfc1bbff9dab5fb4aa038b0) ([#7332](https://github.com/yt-dlp/yt-dlp/issues/7332)) by [garret1317](https://github.com/garret1317)
- **niconico**
    - [Download comments from the new endpoint](https://github.com/yt-dlp/yt-dlp/commit/52ecc33e221f7de7eb6fed6c22489f0c5fdd2c6d) ([#6773](https://github.com/yt-dlp/yt-dlp/issues/6773)) by [Lesmiscore](https://github.com/Lesmiscore)
    - live: [Add extractor](https://github.com/yt-dlp/yt-dlp/commit/f8f9250fe280d37f0988646cd5cc0072f4d33a6d) ([#5764](https://github.com/yt-dlp/yt-dlp/issues/5764)) by [Lesmiscore](https://github.com/Lesmiscore)
    - series: [Fix extraction](https://github.com/yt-dlp/yt-dlp/commit/c86e433c35fe5da6cb29f3539eef97497f84ed38) ([#6898](https://github.com/yt-dlp/yt-dlp/issues/6898)) by [sqrtNOT](https://github.com/sqrtNOT)
- **nubilesporn**: [Add extractor](https://github.com/yt-dlp/yt-dlp/commit/d4e6ef40772e0560a8ed33b844ef7549e86837be) ([#6231](https://github.com/yt-dlp/yt-dlp/issues/6231)) by [permunkle](https://github.com/permunkle)
- **odnoklassniki**: [Fix formats extraction](https://github.com/yt-dlp/yt-dlp/commit/1a2eb5bda51d8b7a78a65acebf72a0dcf9da196b) ([#7217](https://github.com/yt-dlp/yt-dlp/issues/7217)) by [bashonly](https://github.com/bashonly)
- **opencast**
    - [Add ltitools to `_VALID_URL`](https://github.com/yt-dlp/yt-dlp/commit/3588be59cee429a0ab5c4ceb2f162298bb44147d) ([#6371](https://github.com/yt-dlp/yt-dlp/issues/6371)) by [C0D3D3V](https://github.com/C0D3D3V)
    - [Fix format bug](https://github.com/yt-dlp/yt-dlp/commit/89dbf0848370deaa55af88c3593a2a264124caf5) ([#6512](https://github.com/yt-dlp/yt-dlp/issues/6512)) by [C0D3D3V](https://github.com/C0D3D3V)
- **owncloud**: [Add extractor](https://github.com/yt-dlp/yt-dlp/commit/c6d4b82a8b8bce59b1c9ce5e6d349ea428dac0a7) ([#6533](https://github.com/yt-dlp/yt-dlp/issues/6533)) by [C0D3D3V](https://github.com/C0D3D3V)
- **Parler**: [Rewrite extractor](https://github.com/yt-dlp/yt-dlp/commit/80ea6d3dea8483cddd39fc89b5ee1fc06670c33c) ([#6446](https://github.com/yt-dlp/yt-dlp/issues/6446)) by [JChris246](https://github.com/JChris246)
- **pgatour**: [Add extractor](https://github.com/yt-dlp/yt-dlp/commit/3ae182ad89e1427ff7b1684d6a44ff93fa857a0c) ([#6613](https://github.com/yt-dlp/yt-dlp/issues/6613)) by [bashonly](https://github.com/bashonly)
- **playsuisse**: [Support new url format](https://github.com/yt-dlp/yt-dlp/commit/94627c5dde12a72766bdba36e056916c29c40ed1) ([#6528](https://github.com/yt-dlp/yt-dlp/issues/6528)) by [sbor23](https://github.com/sbor23)
- **polskieradio**: [Improve extractors](https://github.com/yt-dlp/yt-dlp/commit/738c90a463257634455ada3e5c18b714c531dede) ([#5948](https://github.com/yt-dlp/yt-dlp/issues/5948)) by [selfisekai](https://github.com/selfisekai)
- **pornez**: [Support new URL formats](https://github.com/yt-dlp/yt-dlp/commit/cbdf9408e6f1e35e98fd6477b3d6902df5b8a47f) ([#6792](https://github.com/yt-dlp/yt-dlp/issues/6792)) by [zhgwn](https://github.com/zhgwn)
- **pornhub**: [Set access cookies to fix extraction](https://github.com/yt-dlp/yt-dlp/commit/62beefa818c75c20b6941389bb197051554a5d41) ([#6685](https://github.com/yt-dlp/yt-dlp/issues/6685)) by [arobase-che](https://github.com/arobase-che), [Schmoaaaaah](https://github.com/Schmoaaaaah)
- **rai**: [Rewrite extractors](https://github.com/yt-dlp/yt-dlp/commit/c6d3f81a4077aaf9cffc6aa2d0dec92f38e74bb0) ([#5940](https://github.com/yt-dlp/yt-dlp/issues/5940)) by [danog](https://github.com/danog), [nixxo](https://github.com/nixxo)
- **recurbate**: [Add extractor](https://github.com/yt-dlp/yt-dlp/commit/c2502cfed91415c7ccfff925fd3404d230046484) ([#6297](https://github.com/yt-dlp/yt-dlp/issues/6297)) by [mrscrapy](https://github.com/mrscrapy)
- **reddit**
    - [Add login support](https://github.com/yt-dlp/yt-dlp/commit/4d9280c9c853733534dda60486fa949bcca36c9e) ([#6950](https://github.com/yt-dlp/yt-dlp/issues/6950)) by [bashonly](https://github.com/bashonly)
    - [Support cookies and short URLs](https://github.com/yt-dlp/yt-dlp/commit/7a6f6f24592a8065376f11a58e44878807732cf6) ([#6825](https://github.com/yt-dlp/yt-dlp/issues/6825)) by [bashonly](https://github.com/bashonly)
- **rokfin**: [Re-construct manifest url](https://github.com/yt-dlp/yt-dlp/commit/7a6c8a0807941dd24fbf0d6172e811884f98e027) ([#6507](https://github.com/yt-dlp/yt-dlp/issues/6507)) by [vampirefrog](https://github.com/vampirefrog)
- **rottentomatoes**: [Fix extractor](https://github.com/yt-dlp/yt-dlp/commit/2d306c03d6f2697fcbabb7da35aa62cc078359d3) ([#6844](https://github.com/yt-dlp/yt-dlp/issues/6844)) by [JChris246](https://github.com/JChris246)
- **rozhlas**
    - [Extract manifest formats](https://github.com/yt-dlp/yt-dlp/commit/e4cf7741f9302b3faa092962f2895b55cb3d89bb) ([#6590](https://github.com/yt-dlp/yt-dlp/issues/6590)) by [bashonly](https://github.com/bashonly)
    - `MujRozhlas`: [Add extractor](https://github.com/yt-dlp/yt-dlp/commit/c2b801fea59628d5c873e06a0727fbf2051bbd1f) ([#7129](https://github.com/yt-dlp/yt-dlp/issues/7129)) by [stanoarn](https://github.com/stanoarn)
- **rtvc**: [Add extractors](https://github.com/yt-dlp/yt-dlp/commit/9b30cd3dfce83c2f0201b28a7a3ef44ab9722664) ([#6578](https://github.com/yt-dlp/yt-dlp/issues/6578)) by [elyse0](https://github.com/elyse0)
- **rumble**
    - [Detect timeline format](https://github.com/yt-dlp/yt-dlp/commit/78bc1868ff3352108ab2911033d1ac67a55f151e) by [pukkandan](https://github.com/pukkandan)
    - [Fix videos without quality selection](https://github.com/yt-dlp/yt-dlp/commit/6994afc030d2a786d8032075ed71a14d7eac5a4f) by [pukkandan](https://github.com/pukkandan)
- **sbs**: [Overhaul extractor for new API](https://github.com/yt-dlp/yt-dlp/commit/6a765f135ccb654861336ea27a2c1c24ea8e286f) ([#6839](https://github.com/yt-dlp/yt-dlp/issues/6839)) by [bashonly](https://github.com/bashonly), [dirkf](https://github.com/dirkf), [vidiot720](https://github.com/vidiot720)
- **shemaroome**: [Pass `stream_key` header to downloader](https://github.com/yt-dlp/yt-dlp/commit/7bc92517463f5766e9d9b92c3823b5cf403c0e3d) ([#7224](https://github.com/yt-dlp/yt-dlp/issues/7224)) by [bashonly](https://github.com/bashonly)
- **sonyliv**: [Fix login with token](https://github.com/yt-dlp/yt-dlp/commit/4815d35c191e7d375b94492a6486dd2ba43a8954) ([#7223](https://github.com/yt-dlp/yt-dlp/issues/7223)) by [bashonly](https://github.com/bashonly)
- **stageplus**: [Add extractor](https://github.com/yt-dlp/yt-dlp/commit/e5265dc6517478e589ee3c1ff0cb19bdf4e35ce1) ([#6838](https://github.com/yt-dlp/yt-dlp/issues/6838)) by [bashonly](https://github.com/bashonly)
- **stripchat**: [Fix extractor](https://github.com/yt-dlp/yt-dlp/commit/f9213f8a2d7ba46b912afe1dd3ce6bb700a33d72) ([#7306](https://github.com/yt-dlp/yt-dlp/issues/7306)) by [foreignBlade](https://github.com/foreignBlade)
- **substack**: [Fix extraction](https://github.com/yt-dlp/yt-dlp/commit/12037d8b0a578fcc78a5c8f98964e48ee6060e25) ([#7218](https://github.com/yt-dlp/yt-dlp/issues/7218)) by [bashonly](https://github.com/bashonly)
- **sverigesradio**: [Support slug URLs](https://github.com/yt-dlp/yt-dlp/commit/5ee9a7d6e18ceea956e831994cf11c423979354f) ([#7220](https://github.com/yt-dlp/yt-dlp/issues/7220)) by [bashonly](https://github.com/bashonly)
- **tagesschau**: [Fix single audio urls](https://github.com/yt-dlp/yt-dlp/commit/af7585c824a1e405bd8afa46d87b4be322edc93c) ([#6626](https://github.com/yt-dlp/yt-dlp/issues/6626)) by [flashdagger](https://github.com/flashdagger)
- **teamcoco**: [Fix extractor](https://github.com/yt-dlp/yt-dlp/commit/c459d45dd4d417fb80a52e1a04e607776a44baa4) ([#6437](https://github.com/yt-dlp/yt-dlp/issues/6437)) by [bashonly](https://github.com/bashonly)
- **telecaribe**: [Expand livestream support](https://github.com/yt-dlp/yt-dlp/commit/69b2f838d3d3e37dc17367ef64d978db1bea45cf) ([#6601](https://github.com/yt-dlp/yt-dlp/issues/6601)) by [bashonly](https://github.com/bashonly)
- **tencent**: [Fix fatal metadata extraction](https://github.com/yt-dlp/yt-dlp/commit/971d901d129403e875a04dd92109507a03fbc070) ([#7219](https://github.com/yt-dlp/yt-dlp/issues/7219)) by [bashonly](https://github.com/bashonly)
- **thesun**: [Update `_VALID_URL`](https://github.com/yt-dlp/yt-dlp/commit/0181b9a1b31db3fde943f7cd3fe9662f23bff292) ([#6522](https://github.com/yt-dlp/yt-dlp/issues/6522)) by [hatienl0i261299](https://github.com/hatienl0i261299)
- **tiktok**
    - [Extract 1080p adaptive formats](https://github.com/yt-dlp/yt-dlp/commit/c2a1bdb00931969193f2a31ea27b9c66a07aaec2) ([#7228](https://github.com/yt-dlp/yt-dlp/issues/7228)) by [bashonly](https://github.com/bashonly)
    - [Fix and improve metadata extraction](https://github.com/yt-dlp/yt-dlp/commit/925936908a3c3ee0e508621db14696b9f6a8b563) ([#6777](https://github.com/yt-dlp/yt-dlp/issues/6777)) by [bashonly](https://github.com/bashonly)
    - [Fix mp3 formats](https://github.com/yt-dlp/yt-dlp/commit/8ceb07e870424c219dced8f4348729553f05c5cc) ([#6615](https://github.com/yt-dlp/yt-dlp/issues/6615)) by [bashonly](https://github.com/bashonly)
    - [Fix resolution extraction](https://github.com/yt-dlp/yt-dlp/commit/ab6057ec80aa75db6303b8206916d00c376c622c) ([#7237](https://github.com/yt-dlp/yt-dlp/issues/7237)) by [puc9](https://github.com/puc9)
    - [Improve `TikTokLive` extractor](https://github.com/yt-dlp/yt-dlp/commit/216bcb66d7dce0762767d751dad10650cb57da9d) ([#6520](https://github.com/yt-dlp/yt-dlp/issues/6520)) by [bashonly](https://github.com/bashonly)
- **triller**: [Support short URLs, detect removed videos](https://github.com/yt-dlp/yt-dlp/commit/33b737bedf8383c0d00d4e1d06a5273dcdfdb756) ([#6636](https://github.com/yt-dlp/yt-dlp/issues/6636)) by [bashonly](https://github.com/bashonly)
- **tv4**: [Fix extractor](https://github.com/yt-dlp/yt-dlp/commit/125ffaa1737dd04716f2f6fbb0595ad3eb7a4b1c) ([#5649](https://github.com/yt-dlp/yt-dlp/issues/5649)) by [dirkf](https://github.com/dirkf), [TxI5](https://github.com/TxI5)
- **tvp**: [Use new API](https://github.com/yt-dlp/yt-dlp/commit/0c7ce146e4d2a84e656d78f6857952bfd25ab389) ([#6989](https://github.com/yt-dlp/yt-dlp/issues/6989)) by [selfisekai](https://github.com/selfisekai)
- **tvplay**: [Remove outdated domains](https://github.com/yt-dlp/yt-dlp/commit/937264419f9bf375d5656785ae6e53282587c15d) ([#7106](https://github.com/yt-dlp/yt-dlp/issues/7106)) by [ivanskodje](https://github.com/ivanskodje)
- **twitch**
    - [Extract original size thumbnail](https://github.com/yt-dlp/yt-dlp/commit/80b732b7a9585b2a61e456dc0d2d014a439cbaee) ([#6629](https://github.com/yt-dlp/yt-dlp/issues/6629)) by [JC-Chung](https://github.com/JC-Chung)
    - [Fix `is_live`](https://github.com/yt-dlp/yt-dlp/commit/0551511b45f7847f40e4314aa9e624e80d086539) ([#6500](https://github.com/yt-dlp/yt-dlp/issues/6500)) by [elyse0](https://github.com/elyse0)
    - [Support mobile clips](https://github.com/yt-dlp/yt-dlp/commit/02312c03cf53eb1da24c9ad022ee79af26060733) ([#6699](https://github.com/yt-dlp/yt-dlp/issues/6699)) by [bepvte](https://github.com/bepvte)
    - [Update `_CLIENT_ID` and add extractor-arg](https://github.com/yt-dlp/yt-dlp/commit/01231feb142e80828985aabdec04ac608e3d43e2) ([#7200](https://github.com/yt-dlp/yt-dlp/issues/7200)) by [bashonly](https://github.com/bashonly)
    - vod: [Support links from schedule tab](https://github.com/yt-dlp/yt-dlp/commit/dbce5afa6bb61f6272ade613f2e9a3d66b88c7ea) ([#7071](https://github.com/yt-dlp/yt-dlp/issues/7071)) by [falbrechtskirchinger](https://github.com/falbrechtskirchinger)
- **twitter**
    - [Add login support](https://github.com/yt-dlp/yt-dlp/commit/d1795f4a6af99c976c9d3ea2dabe5cf4f8965d3c) ([#7258](https://github.com/yt-dlp/yt-dlp/issues/7258)) by [bashonly](https://github.com/bashonly)
    - [Default to GraphQL, handle auth errors](https://github.com/yt-dlp/yt-dlp/commit/147e62fc584c3ea6fdb09bb7a47905df68553a22) ([#6957](https://github.com/yt-dlp/yt-dlp/issues/6957)) by [bashonly](https://github.com/bashonly)
    - spaces: [Add `release_timestamp`](https://github.com/yt-dlp/yt-dlp/commit/1c16d9df5330819cc79ad588b24aa5b72765c168) ([#7186](https://github.com/yt-dlp/yt-dlp/issues/7186)) by [CeruleanSky](https://github.com/CeruleanSky)
- **urplay**: [Extract all subtitles](https://github.com/yt-dlp/yt-dlp/commit/7bcd4813215ac98daa4949af2ffc677c78307a38) ([#7309](https://github.com/yt-dlp/yt-dlp/issues/7309)) by [hoaluvn](https://github.com/hoaluvn)
- **voot**: [Fix extractor](https://github.com/yt-dlp/yt-dlp/commit/4f7b11cc1c1cebf598107e00cd7295588ed484da) ([#7227](https://github.com/yt-dlp/yt-dlp/issues/7227)) by [bashonly](https://github.com/bashonly)
- **vrt**: [Overhaul extractors](https://github.com/yt-dlp/yt-dlp/commit/1a7dcca378e80a387923ee05c250d8ba122441c6) ([#6244](https://github.com/yt-dlp/yt-dlp/issues/6244)) by [bashonly](https://github.com/bashonly), [bergoid](https://github.com/bergoid), [jeroenj](https://github.com/jeroenj)
- **weverse**: [Add extractors](https://github.com/yt-dlp/yt-dlp/commit/b844a3f8b16500663e7ab6c6ec061cc9b30f71ac) ([#6711](https://github.com/yt-dlp/yt-dlp/issues/6711)) by [bashonly](https://github.com/bashonly) (With fixes in [fd5d93f](https://github.com/yt-dlp/yt-dlp/commit/fd5d93f7040f9776fd541f4e4079dad7d3b3fb4f))
- **wevidi**: [Add extractor](https://github.com/yt-dlp/yt-dlp/commit/1ea15603d852971ed7d92f4de12808b27b3d9370) ([#6868](https://github.com/yt-dlp/yt-dlp/issues/6868)) by [truedread](https://github.com/truedread)
- **weyyak**: [Add extractor](https://github.com/yt-dlp/yt-dlp/commit/6dc00acf0f1f1107a626c21befd1691403e6aeeb) ([#7124](https://github.com/yt-dlp/yt-dlp/issues/7124)) by [ItzMaxTV](https://github.com/ItzMaxTV)
- **whyp**: [Add extractor](https://github.com/yt-dlp/yt-dlp/commit/2c566ed14101673c651c08c306c30fa5b4010b85) ([#6803](https://github.com/yt-dlp/yt-dlp/issues/6803)) by [CoryTibbettsDev](https://github.com/CoryTibbettsDev)
- **wrestleuniverse**
    - [Fix cookies support](https://github.com/yt-dlp/yt-dlp/commit/c8561c6d03f025268d6d3972abeb47987c8d7cbb) by [bashonly](https://github.com/bashonly)
    - [Fix extraction, add login](https://github.com/yt-dlp/yt-dlp/commit/ef8fb7f029b816dfc95600727d84400591a3b5c5) ([#6982](https://github.com/yt-dlp/yt-dlp/issues/6982)) by [bashonly](https://github.com/bashonly), [Grub4K](https://github.com/Grub4K)
- **wykop**: [Add extractors](https://github.com/yt-dlp/yt-dlp/commit/aed945e1b9b7d3af2a907e1a12e6508cc81d6a20) ([#6140](https://github.com/yt-dlp/yt-dlp/issues/6140)) by [selfisekai](https://github.com/selfisekai)
- **ximalaya**: [Sort playlist entries](https://github.com/yt-dlp/yt-dlp/commit/8790ea7b2536332777bce68590386b1aa935fac7) ([#7292](https://github.com/yt-dlp/yt-dlp/issues/7292)) by [linsui](https://github.com/linsui)
- **YahooGyaOIE, YahooGyaOPlayerIE**: [Delete extractors due to website close](https://github.com/yt-dlp/yt-dlp/commit/68be95bd0ca3f76aa63c9812935bd826b3a42e53) ([#6218](https://github.com/yt-dlp/yt-dlp/issues/6218)) by [Lesmiscore](https://github.com/Lesmiscore)
- **yappy**: YappyProfile: [Add extractor](https://github.com/yt-dlp/yt-dlp/commit/6f69101dc912690338d32e2aab085c32e44eba3f) ([#7346](https://github.com/yt-dlp/yt-dlp/issues/7346)) by [7vlad7](https://github.com/7vlad7)
- **youku**: [Improve error message](https://github.com/yt-dlp/yt-dlp/commit/ef0848abd425dfda6db62baa8d72897eefb0007f) ([#6690](https://github.com/yt-dlp/yt-dlp/issues/6690)) by [carusocr](https://github.com/carusocr)
- **youporn**: [Extract m3u8 formats](https://github.com/yt-dlp/yt-dlp/commit/ddae33754ae1f32dd9c64cf895c47d20f6b5f336) by [pukkandan](https://github.com/pukkandan)
- **youtube**
    - [Add client name to `format_note` when `-v`](https://github.com/yt-dlp/yt-dlp/commit/c795c39f27244cbce846067891827e4847036441) ([#6254](https://github.com/yt-dlp/yt-dlp/issues/6254)) by [Lesmiscore](https://github.com/Lesmiscore), [pukkandan](https://github.com/pukkandan)
    - [Add extractor-arg `include_duplicate_formats`](https://github.com/yt-dlp/yt-dlp/commit/86cb922118b236306310a72657f70426c20e28bb) by [pukkandan](https://github.com/pukkandan)
    - [Bypass throttling for `-f17`](https://github.com/yt-dlp/yt-dlp/commit/c9abebb851e6188cb34b9eb744c1863dd46af919) by [pukkandan](https://github.com/pukkandan)
    - [Construct fragment list lazily](https://github.com/yt-dlp/yt-dlp/commit/2a23d92d9ec44a0168079e38bcf3d383e5c4c7bb) by [pukkandan](https://github.com/pukkandan) (With fixes in [e389d17](https://github.com/yt-dlp/yt-dlp/commit/e389d172b6f42e4f332ae679dc48543fb7b9b61d))
    - [Define strict uploader metadata mapping](https://github.com/yt-dlp/yt-dlp/commit/7666b93604b97e9ada981c6b04ccf5605dd1bd44) ([#6384](https://github.com/yt-dlp/yt-dlp/issues/6384)) by [coletdjnz](https://github.com/coletdjnz)
    - [Determine audio language using automatic captions](https://github.com/yt-dlp/yt-dlp/commit/ff9b0e071ffae5543cc309e6f9e647ac51e5846e) by [pukkandan](https://github.com/pukkandan)
    - [Extract `channel_is_verified`](https://github.com/yt-dlp/yt-dlp/commit/8213ce28a485e200f6a7e1af1434a987c8e702bd) ([#7213](https://github.com/yt-dlp/yt-dlp/issues/7213)) by [coletdjnz](https://github.com/coletdjnz)
    - [Extract `heatmap` data](https://github.com/yt-dlp/yt-dlp/commit/5caf30dbc34f10b0be60676fece635b5c59f0d72) ([#7100](https://github.com/yt-dlp/yt-dlp/issues/7100)) by [tntmod54321](https://github.com/tntmod54321)
    - [Extract more metadata for comments](https://github.com/yt-dlp/yt-dlp/commit/c35448b7b14113b35c4415dbfbf488c4731f006f) ([#7179](https://github.com/yt-dlp/yt-dlp/issues/7179)) by [coletdjnz](https://github.com/coletdjnz)
    - [Extract uploader metadata for feed/playlist items](https://github.com/yt-dlp/yt-dlp/commit/93e12ed76ef49252dc6869b59d21d0777e5e11af) by [coletdjnz](https://github.com/coletdjnz)
    - [Fix comment loop detection for pinned comments](https://github.com/yt-dlp/yt-dlp/commit/141a8dff98874a426d7fbe772e0a8421bb42656f) ([#6714](https://github.com/yt-dlp/yt-dlp/issues/6714)) by [coletdjnz](https://github.com/coletdjnz)
    - [Fix continuation loop with no comments](https://github.com/yt-dlp/yt-dlp/commit/18f8fba7c89a87f99cc3313a1795848867e84fff) ([#7148](https://github.com/yt-dlp/yt-dlp/issues/7148)) by [coletdjnz](https://github.com/coletdjnz)
    - [Fix parsing `comment_count`](https://github.com/yt-dlp/yt-dlp/commit/071670cbeaa01ddf2cc20a95ae6da25f8f086431) ([#6523](https://github.com/yt-dlp/yt-dlp/issues/6523)) by [nick-cd](https://github.com/nick-cd)
    - [Handle incomplete initial data from watch page](https://github.com/yt-dlp/yt-dlp/commit/607510b9f2f67bfe7d33d74031a5c1fe22a24862) ([#6510](https://github.com/yt-dlp/yt-dlp/issues/6510)) by [coletdjnz](https://github.com/coletdjnz)
    - [Ignore wrong fps of some formats](https://github.com/yt-dlp/yt-dlp/commit/97afb093d4cbe5df889145afa5f9ede4535e93e4) by [pukkandan](https://github.com/pukkandan)
    - [Misc cleanup](https://github.com/yt-dlp/yt-dlp/commit/14a14335b280766fbf5a469ae26836d6c1fe450a) by [coletdjnz](https://github.com/coletdjnz)
    - [Prioritize premium formats](https://github.com/yt-dlp/yt-dlp/commit/51a07b0dca4c079d58311c19b6d1c097c24bb021) by [pukkandan](https://github.com/pukkandan)
    - [Revert default formats to `https`](https://github.com/yt-dlp/yt-dlp/commit/c6786ff3baaf72a5baa4d56d34058e54cbcf8ceb) by [pukkandan](https://github.com/pukkandan)
    - [Support podcasts and releases tabs](https://github.com/yt-dlp/yt-dlp/commit/447afb9eaa65bc677e3245c83e53a8e69c174a3c) by [coletdjnz](https://github.com/coletdjnz)
    - [Support shorter relative time format](https://github.com/yt-dlp/yt-dlp/commit/2fb35f6004c7625f0dd493da4a5abf0690f7777c) ([#7191](https://github.com/yt-dlp/yt-dlp/issues/7191)) by [coletdjnz](https://github.com/coletdjnz)
    - music_search_url: [Extract title](https://github.com/yt-dlp/yt-dlp/commit/69a40e4a7f6caa5662527ebd2f3c4e8aa02857a2) ([#7102](https://github.com/yt-dlp/yt-dlp/issues/7102)) by [kangalio](https://github.com/kangalio)
- **zaiko**
    - [Add extractor](https://github.com/yt-dlp/yt-dlp/commit/345b4c0aedd9d19898ce00d5cef35fe0d277a052) ([#7254](https://github.com/yt-dlp/yt-dlp/issues/7254)) by [c-basalt](https://github.com/c-basalt)
    - ZaikoETicket: [Add extractor](https://github.com/yt-dlp/yt-dlp/commit/5cc09c004bd5edbbada9b041c08a720cadc4f4df) ([#7347](https://github.com/yt-dlp/yt-dlp/issues/7347)) by [pzhlkj6612](https://github.com/pzhlkj6612)
- **zdf**: [Fix formats extraction](https://github.com/yt-dlp/yt-dlp/commit/ee0ed0338df328cd986f97315c8162b5a151476d) by [bashonly](https://github.com/bashonly)
- **zee5**: [Fix extraction of new content](https://github.com/yt-dlp/yt-dlp/commit/9d7fde89a40360396f0baa2ee8bf507f92108b32) ([#7280](https://github.com/yt-dlp/yt-dlp/issues/7280)) by [bashonly](https://github.com/bashonly)
- **zingmp3**: [Fix and improve extractors](https://github.com/yt-dlp/yt-dlp/commit/17d7ca84ea723c20668bd9bfa938be7ea0e64f6b) ([#6367](https://github.com/yt-dlp/yt-dlp/issues/6367)) by [hatienl0i261299](https://github.com/hatienl0i261299)
- **zoom**
    - [Fix extractor](https://github.com/yt-dlp/yt-dlp/commit/79c77e85b70ae3b9942d5a88c14d021a9bd24222) ([#6741](https://github.com/yt-dlp/yt-dlp/issues/6741)) by [shreyasminocha](https://github.com/shreyasminocha)
    - [Fix share URL extraction](https://github.com/yt-dlp/yt-dlp/commit/90c1f5120694105496a6ad9e3ecfc6c25de6cae1) ([#6789](https://github.com/yt-dlp/yt-dlp/issues/6789)) by [bashonly](https://github.com/bashonly)

#### Downloader changes
- **curl**: [Fix progress reporting](https://github.com/yt-dlp/yt-dlp/commit/66aeaac9aa30b5959069ba84e53a5508232deb38) by [pukkandan](https://github.com/pukkandan)
- **fragment**: [Do not sleep between fragments](https://github.com/yt-dlp/yt-dlp/commit/424f3bf03305088df6e01d62f7311be8601ad3f4) by [pukkandan](https://github.com/pukkandan)

#### Postprocessor changes
- [Fix chapters if duration is not extracted](https://github.com/yt-dlp/yt-dlp/commit/01ddec7e661bf90dc4c34e6924eb9d7629886cef) ([#6037](https://github.com/yt-dlp/yt-dlp/issues/6037)) by [bashonly](https://github.com/bashonly)
- [Print newline for `--progress-template`](https://github.com/yt-dlp/yt-dlp/commit/13ff78095372fd98900a32572cf817994c07ccb5) by [pukkandan](https://github.com/pukkandan)
- **EmbedThumbnail, FFmpegMetadata**: [Fix error on attaching thumbnails and info json for mkv/mka](https://github.com/yt-dlp/yt-dlp/commit/0f0875ed555514f32522a0f30554fb08825d5124) ([#6647](https://github.com/yt-dlp/yt-dlp/issues/6647)) by [Lesmiscore](https://github.com/Lesmiscore)
- **FFmpegFixupM3u8PP**: [Check audio codec before fixup](https://github.com/yt-dlp/yt-dlp/commit/3f7e2bd80e3c5d8a1682f20a1b245fcd974f295d) ([#6778](https://github.com/yt-dlp/yt-dlp/issues/6778)) by [bashonly](https://github.com/bashonly)
- **FixupDuplicateMoov**: [Fix bug in triggering](https://github.com/yt-dlp/yt-dlp/commit/26010b5cec50193b98ad7845d1d77450f9f14c2b) by [pukkandan](https://github.com/pukkandan)

#### Misc. changes
- [Add automatic duplicate issue detection](https://github.com/yt-dlp/yt-dlp/commit/15b2d3db1d40b0437fca79d8874d392aa54b3cdd) by [pukkandan](https://github.com/pukkandan)
- **build**
    - [Fix macOS target](https://github.com/yt-dlp/yt-dlp/commit/44a79958f0b596ee71e1eb25f158610aada29d1b) by [Grub4K](https://github.com/Grub4K)
    - [Implement build verification using `--update-to`](https://github.com/yt-dlp/yt-dlp/commit/b73193c99aa23b135732408a5fcf655c68d731c6) by [bashonly](https://github.com/bashonly), [Grub4K](https://github.com/Grub4K)
    - [Pin `pyinstaller` version for MacOS](https://github.com/yt-dlp/yt-dlp/commit/427a8fafbb0e18c28d0ed7960be838d7b26b88d3) by [pukkandan](https://github.com/pukkandan)
    - [Various build workflow improvements](https://github.com/yt-dlp/yt-dlp/commit/c4efa0aefec8daef1de62fd1693f13edf3c8b03c) by [bashonly](https://github.com/bashonly), [Grub4K](https://github.com/Grub4K)
- **cleanup**
    - Miscellaneous
        - [6f2287c](https://github.com/yt-dlp/yt-dlp/commit/6f2287cb18cbfb27518f068d868fa9390fee78ad) by [pukkandan](https://github.com/pukkandan)
        - [ad54c91](https://github.com/yt-dlp/yt-dlp/commit/ad54c9130e793ce433bf9da334fa80df9f3aee58) by [freezboltz](https://github.com/freezboltz), [mikf](https://github.com/mikf), [pukkandan](https://github.com/pukkandan)
- **cleanup, utils**: [Split into submodules](https://github.com/yt-dlp/yt-dlp/commit/69bec6730ec9d724bcedeab199d9d684d61423ba) ([#7090](https://github.com/yt-dlp/yt-dlp/issues/7090)) by [coletdjnz](https://github.com/coletdjnz), [pukkandan](https://github.com/pukkandan)
- **cli_to_api**: [Add script](https://github.com/yt-dlp/yt-dlp/commit/46f1370e9af6f8af8762f67e27e5acb8f0c48a47) by [pukkandan](https://github.com/pukkandan)
- **devscripts**: `make_changelog`: [Various improvements](https://github.com/yt-dlp/yt-dlp/commit/23c39a4beadee382060bb47fdaa21316ca707d38) by [Grub4K](https://github.com/Grub4K)
- **docs**: [Misc improvements](https://github.com/yt-dlp/yt-dlp/commit/c8bc203fbf3bb09914e53f0833eed622ab7edbb9) by [pukkandan](https://github.com/pukkandan)

### 2023.03.04

#### Extractor changes
- bilibili
    - [Fix for downloading wrong subtitles](https://github.com/yt-dlp/yt-dlp/commit/8a83baaf218ab89e6e7faa76b7c7be3a2ec19e3a) ([#6358](https://github.com/yt-dlp/yt-dlp/issues/6358)) by [LXYan2333](https://github.com/LXYan2333)
- ESPNcricinfo
    - [Handle new URL pattern](https://github.com/yt-dlp/yt-dlp/commit/640c934823fc2d1ec77ec932566078014058635f) ([#6321](https://github.com/yt-dlp/yt-dlp/issues/6321)) by [venkata-krishnas](https://github.com/venkata-krishnas)
- lefigaro
    - [Add extractors](https://github.com/yt-dlp/yt-dlp/commit/eb8fd6d044e8926532772b72be0645c6b8ecb3aa) ([#6309](https://github.com/yt-dlp/yt-dlp/issues/6309)) by [elyse0](https://github.com/elyse0)
- lumni
    - [Add extractor](https://github.com/yt-dlp/yt-dlp/commit/1f8489cccbdc6e96027ef527b88717458f0900e8) ([#6302](https://github.com/yt-dlp/yt-dlp/issues/6302)) by [carusocr](https://github.com/carusocr)
- Prankcast
    - [Fix tags](https://github.com/yt-dlp/yt-dlp/commit/ed4cc4ea793314c50ae3f82e98248c1de1c25694) ([#6316](https://github.com/yt-dlp/yt-dlp/issues/6316)) by [columndeeply](https://github.com/columndeeply)
- rutube
    - [Extract chapters from description](https://github.com/yt-dlp/yt-dlp/commit/22ccd5420b3eb0782776071f12cccd1fedaa1fd0) ([#6345](https://github.com/yt-dlp/yt-dlp/issues/6345)) by [mushbite](https://github.com/mushbite)
- SportDeutschland
    - [Rewrite extractor](https://github.com/yt-dlp/yt-dlp/commit/45db357289b4e1eec09093c8bc5446520378f426) by [pukkandan](https://github.com/pukkandan)
- telecaribe
    - [Add extractor](https://github.com/yt-dlp/yt-dlp/commit/b40471282286bd2b09c485bf79afd271d229272c) ([#6311](https://github.com/yt-dlp/yt-dlp/issues/6311)) by [elyse0](https://github.com/elyse0)
- tubetugraz
    - [Support `--twofactor` (#6424)](https://github.com/yt-dlp/yt-dlp/commit/f44cb4e77bb9be8be291d02ab6f79dc0b4c0d4a1) ([#6427](https://github.com/yt-dlp/yt-dlp/issues/6427)) by [Ferdi265](https://github.com/Ferdi265)
- tunein
    - [Fix extractors](https://github.com/yt-dlp/yt-dlp/commit/46580ced56c90b559885aded6aa8f46f20a9cdce) ([#6310](https://github.com/yt-dlp/yt-dlp/issues/6310)) by [elyse0](https://github.com/elyse0)
- twitch
    - [Update for GraphQL API changes](https://github.com/yt-dlp/yt-dlp/commit/4a6272c6d1bff89969b67cd22b26ebe6d7e72279) ([#6318](https://github.com/yt-dlp/yt-dlp/issues/6318)) by [elyse0](https://github.com/elyse0)
- twitter
    - [Fix retweet extraction](https://github.com/yt-dlp/yt-dlp/commit/cf605226521e99c89fc8dff26a319025810e63a0) ([#6422](https://github.com/yt-dlp/yt-dlp/issues/6422)) by [selfisekai](https://github.com/selfisekai)
- xvideos
    - quickies: [Add extractor](https://github.com/yt-dlp/yt-dlp/commit/283a0b5bc511f3b350eead4488158f50c20ec526) ([#6414](https://github.com/yt-dlp/yt-dlp/issues/6414)) by [Yakabuff](https://github.com/Yakabuff)

#### Misc. changes
- build
    - [Fix publishing to PyPI and homebrew](https://github.com/yt-dlp/yt-dlp/commit/55676fe498345a389a2539d8baaba958d6d61c3e) by [bashonly](https://github.com/bashonly)
    - [Only archive if `vars.ARCHIVE_REPO` is set](https://github.com/yt-dlp/yt-dlp/commit/08ff6d59f97b5f5f0128f6bf6fbef56fd836cc52) by [Grub4K](https://github.com/Grub4K)
- cleanup
    - Miscellaneous: [392389b](https://github.com/yt-dlp/yt-dlp/commit/392389b7df7b818f794b231f14dc396d4875fbad) by [pukkandan](https://github.com/pukkandan)
- devscripts
    - `make_changelog`: [Stop at `Release ...` commit](https://github.com/yt-dlp/yt-dlp/commit/7accdd9845fe7ce9d0aa5a9d16faaa489c1294eb) by [pukkandan](https://github.com/pukkandan)

### 2023.03.03

#### Important changes
- **A new release type has been added!**
    * [`nightly`](https://github.com/yt-dlp/yt-dlp/releases/tag/nightly) builds will be made after each push, containing the latest fixes (but also possibly bugs).
    * When using `--update`/`-U`, a release binary will only update to its current channel (either `stable` or `nightly`).
    * The `--update-to` option has been added allowing the user more control over program upgrades (or downgrades).
    * `--update-to` can change the release channel (`stable`, `nightly`) and also upgrade or downgrade to specific tags.
    * **Usage**: `--update-to CHANNEL`, `--update-to TAG`, `--update-to CHANNEL@TAG`
- **YouTube throttling fixes!**

#### Core changes
- [Add option `--break-match-filters`](https://github.com/yt-dlp/yt-dlp/commit/fe2ce85aff0aa03735fc0152bb8cb9c3d4ef0753) by [pukkandan](https://github.com/pukkandan)
- [Fix `--break-on-existing` with `--lazy-playlist`](https://github.com/yt-dlp/yt-dlp/commit/d21056f4cf0a1623daa107f9181074f5725ac436) by [pukkandan](https://github.com/pukkandan)
- dependencies
    - [Simplify `Cryptodome`](https://github.com/yt-dlp/yt-dlp/commit/65f6e807804d2af5e00f2aecd72bfc43af19324a) by [pukkandan](https://github.com/pukkandan)
- jsinterp
    - [Handle `Date` at epoch 0](https://github.com/yt-dlp/yt-dlp/commit/9acf1ee25f7ad3920ede574a9de95b8c18626af4) by [pukkandan](https://github.com/pukkandan)
- plugins
    - [Don't look in `.egg` directories](https://github.com/yt-dlp/yt-dlp/commit/b059188383eee4fa336ef728dda3ff4bb7335625) by [pukkandan](https://github.com/pukkandan)
- update
    - [Add option `--update-to`, including to nightly](https://github.com/yt-dlp/yt-dlp/commit/77df20f14cc9ed41dfe3a1fe2d77fd27f5365a94) ([#6220](https://github.com/yt-dlp/yt-dlp/issues/6220)) by [bashonly](https://github.com/bashonly), [Grub4K](https://github.com/Grub4K), [pukkandan](https://github.com/pukkandan)
- utils
    - `LenientJSONDecoder`: [Parse unclosed objects](https://github.com/yt-dlp/yt-dlp/commit/cc09083636ce21e58ff74f45eac2dbda507462b0) by [pukkandan](https://github.com/pukkandan)
    - `Popen`: [Shim undocumented `text_mode` property](https://github.com/yt-dlp/yt-dlp/commit/da8e2912b165005f76779a115a071cd6132ceedf) by [Grub4K](https://github.com/Grub4K)

#### Extractor changes
- [Fix DRM detection in m3u8](https://github.com/yt-dlp/yt-dlp/commit/43a3eaf96393b712d60cbcf5c6cb1e90ed7f42f5) by [pukkandan](https://github.com/pukkandan)
- generic
    - [Detect manifest links via extension](https://github.com/yt-dlp/yt-dlp/commit/b38cae49e6f4849c8ee2a774bdc3c1c647ae5f0e) by [bashonly](https://github.com/bashonly)
    - [Handle basic-auth when checking redirects](https://github.com/yt-dlp/yt-dlp/commit/8e9fe43cd393e69fa49b3d842aa3180c1d105b8f) by [pukkandan](https://github.com/pukkandan)
- GoogleDrive
    - [Fix some audio](https://github.com/yt-dlp/yt-dlp/commit/4d248e29d20d983ededab0b03d4fe69dff9eb4ed) by [pukkandan](https://github.com/pukkandan)
- iprima
    - [Fix extractor](https://github.com/yt-dlp/yt-dlp/commit/9fddc12ab022a31754e0eaa358fc4e1dfa974587) ([#6291](https://github.com/yt-dlp/yt-dlp/issues/6291)) by [std-move](https://github.com/std-move)
- mediastream
    - [Improve WinSports support](https://github.com/yt-dlp/yt-dlp/commit/2d5a8c5db2bd4ff1c2e45e00cd890a10f8ffca9e) ([#6401](https://github.com/yt-dlp/yt-dlp/issues/6401)) by [bashonly](https://github.com/bashonly)
- ntvru
    - [Extract HLS and DASH formats](https://github.com/yt-dlp/yt-dlp/commit/77d6d136468d0c23c8e79bc937898747804f585a) ([#6403](https://github.com/yt-dlp/yt-dlp/issues/6403)) by [bashonly](https://github.com/bashonly)
- tencent
    - [Add more formats and info](https://github.com/yt-dlp/yt-dlp/commit/18d295c9e0f95adc179eef345b7af64d6372db78) ([#5950](https://github.com/yt-dlp/yt-dlp/issues/5950)) by [Hill-98](https://github.com/Hill-98)
- yle_areena
    - [Extract non-Kaltura videos](https://github.com/yt-dlp/yt-dlp/commit/40d77d89027cd0e0ce31d22aec81db3e1d433900) ([#6402](https://github.com/yt-dlp/yt-dlp/issues/6402)) by [bashonly](https://github.com/bashonly)
- youtube
    - [Construct dash formats with `range` query](https://github.com/yt-dlp/yt-dlp/commit/5038f6d713303e0967d002216e7a88652401c22a) by [pukkandan](https://github.com/pukkandan) (With fixes in [f34804b](https://github.com/yt-dlp/yt-dlp/commit/f34804b2f920f62a6e893a14a9e2a2144b14dd23) by [bashonly](https://github.com/bashonly), [coletdjnz](https://github.com/coletdjnz))
    - [Detect and break on looping comments](https://github.com/yt-dlp/yt-dlp/commit/7f51861b1820c37b157a239b1fe30628d907c034) ([#6301](https://github.com/yt-dlp/yt-dlp/issues/6301)) by [coletdjnz](https://github.com/coletdjnz)
    - [Extract channel `view_count` when `/about` tab is passed](https://github.com/yt-dlp/yt-dlp/commit/31e183557fcd1b937582f9429f29207c1261f501) by [pukkandan](https://github.com/pukkandan)

#### Misc. changes
- build
    - [Add `cffi` as a dependency for `yt_dlp_linux`](https://github.com/yt-dlp/yt-dlp/commit/776d1c3f0c9b00399896dd2e40e78e9a43218109) by [bashonly](https://github.com/bashonly)
    - [Automated builds and nightly releases](https://github.com/yt-dlp/yt-dlp/commit/29cb20bd563c02671b31dd840139e93dd37150a1) ([#6220](https://github.com/yt-dlp/yt-dlp/issues/6220)) by [bashonly](https://github.com/bashonly), [Grub4K](https://github.com/Grub4K) (With fixes in [bfc861a](https://github.com/yt-dlp/yt-dlp/commit/bfc861a91ee65c9b0ac169754f512e052c6827cf) by [pukkandan](https://github.com/pukkandan))
    - [Sign SHA files and release public key](https://github.com/yt-dlp/yt-dlp/commit/12647e03d417feaa9ea6a458bea5ebd747494a53) by [Grub4K](https://github.com/Grub4K)
- cleanup
    - [Fix `Changelog`](https://github.com/yt-dlp/yt-dlp/commit/17ca19ab60a6a13eb8a629c51442b5248b0d8394) by [pukkandan](https://github.com/pukkandan)
    - jsinterp: [Give functions names to help debugging](https://github.com/yt-dlp/yt-dlp/commit/b2e0343ba0fc5d8702e90f6ba2b71358e2677e0b) by [pukkandan](https://github.com/pukkandan)
    - Miscellaneous: [4815bbf](https://github.com/yt-dlp/yt-dlp/commit/4815bbfc41cf641e4a0650289dbff968cb3bde76), [5b28cef](https://github.com/yt-dlp/yt-dlp/commit/5b28cef72db3b531680d89c121631c73ae05354f) by [pukkandan](https://github.com/pukkandan)
- devscripts
    - [Script to generate changelog](https://github.com/yt-dlp/yt-dlp/commit/d400e261cf029a3f20d364113b14de973be75404) ([#6220](https://github.com/yt-dlp/yt-dlp/issues/6220)) by [Grub4K](https://github.com/Grub4K) (With fixes in [9344964](https://github.com/yt-dlp/yt-dlp/commit/93449642815a6973a4b09b289982ca7e1f961b5f))

### 2023.02.17

* Merge youtube-dl: Upto [commit/2dd6c6e](https://github.com/ytdl-org/youtube-dl/commit/2dd6c6e)
* Fix `--concat-playlist`
* Imply `--no-progress` when `--print`
* Improve default subtitle language selection by [sdht0](https://github.com/sdht0)
* Make `title` completely non-fatal
* Sanitize formats before sorting by [pukkandan](https://github.com/pukkandan)
* Support module level `__bool__` and `property`
* [dependencies] Standardize `Cryptodome` imports
* [hls] Allow extractors to provide AES key by [Grub4K](https://github.com/Grub4K), [bashonly](https://github.com/bashonly)
* [ExtractAudio] Handle outtmpl without ext by [carusocr](https://github.com/carusocr)
* [extractor/common] Fix `_search_nuxt_data` by [LowSuggestion912](https://github.com/LowSuggestion912)
* [extractor/generic] Avoid catastrophic backtracking in KVS regex by [bashonly](https://github.com/bashonly)
* [jsinterp] Support `if` statements
* [plugins] Fix zip search paths
* [utils] `traverse_obj`:  Various improvements by [Grub4K](https://github.com/Grub4K)
* [utils] `traverse_obj`: Fix more bugs
* [utils] `traverse_obj`: Fix several behavioral problems by [Grub4K](https://github.com/Grub4K)
* [utils] Don't use Content-length with encoding by [felixonmars](https://github.com/felixonmars)
* [utils] Fix `time_seconds` to use the provided TZ by [Grub4K](https://github.com/Grub4K), [Lesmiscore](https://github.com/Lesmiscore)
* [utils] Fix race condition in `make_dir` by [aionescu](https://github.com/aionescu)
* [utils] Use local kernel32 for file locking on Windows by [Grub4K](https://github.com/Grub4K)
* [compat_utils] Improve `passthrough_module`
* [compat_utils] Simplify `EnhancedModule`
* [build] Update pyinstaller
* [pyinst] Fix for pyinstaller 5.8
* [devscripts] Provide `pyinstaller` hooks
* [devscripts/pyinstaller] Analyze sub-modules of `Cryptodome`
* [cleanup] Misc fixes and cleanup
* [extractor/anchorfm] Add episode extractor by [HobbyistDev](https://github.com/HobbyistDev), [bashonly](https://github.com/bashonly)
* [extractor/boxcast] Add extractor by [HobbyistDev](https://github.com/HobbyistDev)
* [extractor/ebay] Add extractor by [JChris246](https://github.com/JChris246)
* [extractor/hypergryph] Add extractor by [HobbyistDev](https://github.com/HobbyistDev), [bashonly](https://github.com/bashonly)
* [extractor/NZOnScreen] Add extractor by [gregsadetsky](https://github.com/gregsadetsky), [pukkandan](https://github.com/pukkandan)
* [extractor/rozhlas] Add extractor RozhlasVltavaIE by [amra](https://github.com/amra)
* [extractor/tempo] Add IVXPlayer extractor by [HobbyistDev](https://github.com/HobbyistDev)
* [extractor/txxx] Add extractors by [chio0hai](https://github.com/chio0hai)
* [extractor/vocaroo] Add extractor by [SuperSonicHub1](https://github.com/SuperSonicHub1), [qbnu](https://github.com/qbnu)
* [extractor/wrestleuniverse] Add extractors by [Grub4K](https://github.com/Grub4K), [bashonly](https://github.com/bashonly)
* [extractor/yappy] Add extractor by [HobbyistDev](https://github.com/HobbyistDev), [dirkf](https://github.com/dirkf)
* [extractor/youtube] **Fix `uploader_id` extraction** by [bashonly](https://github.com/bashonly)
* [extractor/youtube] Add hyperpipe instances by [Generator](https://github.com/Generator)
* [extractor/youtube] Handle `consent.youtube`
* [extractor/youtube] Support `/live/` URL
* [extractor/youtube] Update invidious and piped instances by [rohieb](https://github.com/rohieb)
* [extractor/91porn] Fix title and comment extraction by [pmitchell86](https://github.com/pmitchell86)
* [extractor/AbemaTV] Cache user token whenever appropriate by [Lesmiscore](https://github.com/Lesmiscore)
* [extractor/bfmtv] Support `rmc` prefix by [carusocr](https://github.com/carusocr)
* [extractor/biliintl] Add intro and ending chapters by [HobbyistDev](https://github.com/HobbyistDev)
* [extractor/clyp] Support `wav` by [qulaz](https://github.com/qulaz)
* [extractor/crunchyroll] Add intro chapter by [ByteDream](https://github.com/ByteDream)
* [extractor/crunchyroll] Better message for premium videos
* [extractor/crunchyroll] Fix incorrect premium-only error by [Grub4K](https://github.com/Grub4K)
* [extractor/DouyuTV] Use new API by [hatienl0i261299](https://github.com/hatienl0i261299)
* [extractor/embedly] Embedded links may be for other extractors
* [extractor/freesound] Workaround invalid URL in webpage by [rebane2001](https://github.com/rebane2001)
* [extractor/GoPlay] Use new API by [jeroenj](https://github.com/jeroenj)
* [extractor/Hidive] Fix subtitles and age-restriction by [chexxor](https://github.com/chexxor)
* [extractor/huya] Support HD streams by [felixonmars](https://github.com/felixonmars)
* [extractor/moviepilot] Fix extractor by [panatexxa](https://github.com/panatexxa)
* [extractor/nbc] Fix `NBC` and `NBCStations` extractors by [bashonly](https://github.com/bashonly)
* [extractor/nbc] Fix XML parsing by [bashonly](https://github.com/bashonly)
* [extractor/nebula] Remove broken cookie support by [hheimbuerger](https://github.com/hheimbuerger)
* [extractor/nfl] Add `NFLPlus` extractors by [bashonly](https://github.com/bashonly)
* [extractor/niconico] Add support for like history by [Matumo](https://github.com/Matumo), [pukkandan](https://github.com/pukkandan)
* [extractor/nitter] Update instance list by [OIRNOIR](https://github.com/OIRNOIR)
* [extractor/npo] Fix extractor and add HD support by [seproDev](https://github.com/seproDev)
* [extractor/odkmedia] Add `OnDemandChinaEpisodeIE` by [HobbyistDev](https://github.com/HobbyistDev), [pukkandan](https://github.com/pukkandan)
* [extractor/pornez] Handle relative URLs in iframe by [JChris246](https://github.com/JChris246)
* [extractor/radiko] Fix format sorting for Time Free by [road-master](https://github.com/road-master)
* [extractor/rcs] Fix extractors by [nixxo](https://github.com/nixxo), [pukkandan](https://github.com/pukkandan)
* [extractor/reddit] Support user posts by [OMEGARAZER](https://github.com/OMEGARAZER)
* [extractor/rumble] Fix format sorting by [pukkandan](https://github.com/pukkandan)
* [extractor/servus] Rewrite extractor by [Ashish0804](https://github.com/Ashish0804), [FrankZ85](https://github.com/FrankZ85), [StefanLobbenmeier](https://github.com/StefanLobbenmeier)
* [extractor/slideslive] Fix slides and chapters/duration by [bashonly](https://github.com/bashonly)
* [extractor/SportDeutschland] Fix extractor by [FriedrichRehren](https://github.com/FriedrichRehren)
* [extractor/Stripchat] Fix extractor by [JChris246](https://github.com/JChris246), [bashonly](https://github.com/bashonly)
* [extractor/tnaflix] Fix extractor by [bashonly](https://github.com/bashonly), [oxamun](https://github.com/oxamun)
* [extractor/tvp] Support `stream.tvp.pl` by [selfisekai](https://github.com/selfisekai)
* [extractor/twitter] Fix `--no-playlist` and add media `view_count` when using GraphQL by [Grub4K](https://github.com/Grub4K)
* [extractor/twitter] Fix graphql extraction on some tweets by [selfisekai](https://github.com/selfisekai)
* [extractor/vimeo] Fix `playerConfig` extraction by [LeoniePhiline](https://github.com/LeoniePhiline), [bashonly](https://github.com/bashonly)
* [extractor/viu] Add `ViuOTTIndonesiaIE` extractor by [HobbyistDev](https://github.com/HobbyistDev)
* [extractor/vk] Fix playlists for new API by [the-marenga](https://github.com/the-marenga)
* [extractor/vlive] Replace with `VLiveWebArchiveIE` by [seproDev](https://github.com/seproDev)
* [extractor/ximalaya] Update album `_VALID_URL` by [carusocr](https://github.com/carusocr)
* [extractor/zdf] Use android API endpoint for UHD downloads by [seproDev](https://github.com/seproDev)
* [extractor/drtv] Fix bug in [ab4cbef](https://github.com/yt-dlp/yt-dlp/commit/ab4cbef) by [bashonly](https://github.com/bashonly)


### 2023.01.06

* Fix config locations by [Grub4K](https://github.com/Grub4K), [coletdjnz](https://github.com/coletdjnz), [pukkandan](https://github.com/pukkandan)
* [downloader/aria2c] Disable native progress
* [utils] `mimetype2ext`: `weba` is not standard
* [utils] `windows_enable_vt_mode`: Better error handling
* [build] Add minimal `pyproject.toml`
* [update] Fix updater file removal on windows by [Grub4K](https://github.com/Grub4K)
* [cleanup] Misc fixes and cleanup
* [extractor/aitube] Add extractor by [HobbyistDev](https://github.com/HobbyistDev)
* [extractor/drtv] Add series extractors by [FrederikNS](https://github.com/FrederikNS)
* [extractor/volejtv] Add extractor by [HobbyistDev](https://github.com/HobbyistDev)
* [extractor/xanimu] Add extractor by [JChris246](https://github.com/JChris246)
* [extractor/youtube] Retry manifest refresh for live-from-start by [mzhou](https://github.com/mzhou)
* [extractor/biliintl] Add `/media` to `VALID_URL` by [HobbyistDev](https://github.com/HobbyistDev)
* [extractor/biliIntl] Add fallback to `video_data` by [HobbyistDev](https://github.com/HobbyistDev)
* [extractor/crunchyroll:show] Add `language` to entries by [Chrissi2812](https://github.com/Chrissi2812)
* [extractor/joj] Fix extractor by [OndrejBakan](https://github.com/OndrejBakan), [pukkandan](https://github.com/pukkandan)
* [extractor/nbc] Update graphql query by [jacobtruman](https://github.com/jacobtruman)
* [extractor/reddit] Add subreddit as `channel_id` by [gschizas](https://github.com/gschizas)
* [extractor/tiktok] Add `TikTokLive` extractor by [JC-Chung](https://github.com/JC-Chung)

### 2023.01.02

* **Improve plugin architecture** by [Grub4K](https://github.com/Grub4K), [coletdjnz](https://github.com/coletdjnz), [flashdagger](https://github.com/flashdagger), [pukkandan](https://github.com/pukkandan)
    * Plugins can be loaded in any distribution of yt-dlp (binary, pip, source, etc.) and can be distributed and installed as packages. See [the readme](https://github.com/yt-dlp/yt-dlp/tree/05997b6e98e638d97d409c65bb5eb86da68f3b64#plugins) for more information
* Add `--compat-options 2021,2022`
    * This allows devs to change defaults and make other potentially breaking changes more easily. If you need everything to work exactly as-is, put Use `--compat 2022` in your config to guard against future compat changes.
* [downloader/aria2c] Native progress for aria2c via RPC by [Lesmiscore](https://github.com/Lesmiscore), [pukkandan](https://github.com/pukkandan)
* Merge youtube-dl: Upto [commit/195f22f](https://github.com/ytdl-org/youtube-dl/commit/195f22f6) by [Grub4K](https://github.com/Grub4K), [pukkandan](https://github.com/pukkandan)
* Add pre-processor stage `video`
* Let `--parse/replace-in-metadata` run at any post-processing stage
* Add `--enable-file-urls` by [coletdjnz](https://github.com/coletdjnz)
* Add new field `aspect_ratio`
* Add `ac4` to known codecs
* Add `weba` to known extensions
* [FFmpegVideoConvertor] Add `gif` to `--recode-video`
* Add message when there are no subtitles/thumbnails
* Deprioritize HEVC-over-FLV formats by [Lesmiscore](https://github.com/Lesmiscore)
* Make early reject of `--match-filter` stricter
* Fix `--cookies-from-browser` CLI parsing
* Fix `original_url` in playlists
* Fix bug in writing playlist info-json
* Fix bugs in `PlaylistEntries`
* [downloader/ffmpeg] Fix headers for video+audio formats by [Grub4K](https://github.com/Grub4K), [bashonly](https://github.com/bashonly)
* [extractor] Add a way to distinguish IEs that returns only videos
* [extractor] Implement universal format sorting and deprecate `_sort_formats`
* [extractor] Let `_extract_format` functions obey `--ignore-no-formats`
* [extractor/generic] Add `fragment_query` extractor arg for DASH and HLS by [bashonly](https://github.com/bashonly), [pukkandan](https://github.com/pukkandan)
* [extractor/generic] Decode unicode-escaped embed URLs by [bashonly](https://github.com/bashonly)
* [extractor/generic] Don't report redirect to https
* [extractor/generic] Fix JSON LD manifest extraction by [bashonly](https://github.com/bashonly), [pukkandan](https://github.com/pukkandan)
* [extractor/generic] Use `Accept-Encoding: identity` for initial request by [coletdjnz](https://github.com/coletdjnz)
* [FormatSort] Add `mov` to `vext`
* [jsinterp] Escape regex that looks like nested set
* [webvtt] Handle premature EOF by [flashdagger](https://github.com/flashdagger)
* [utils] `classproperty`: Add cache support
* [utils] `get_exe_version`: Detect broken executables by [dirkf](https://github.com/dirkf), [pukkandan](https://github.com/pukkandan)
* [utils] `js_to_json`: Fix bug in [f55523c](https://github.com/yt-dlp/yt-dlp/commit/f55523c) by [ChillingPepper](https://github.com/ChillingPepper), [pukkandan](https://github.com/pukkandan)
* [utils] Make `ExtractorError` mutable
* [utils] Move `FileDownloader.parse_bytes` into utils
* [utils] Move format sorting code into `utils`
* [utils] `windows_enable_vt_mode`: Proper implementation by [Grub4K](https://github.com/Grub4K)
* [update] Workaround [#5632](https://github.com/yt-dlp/yt-dlp/issues/5632)
* [docs] Improvements
* [cleanup] Misc fixes and cleanup
* [cleanup] Use `random.choices` by [freezboltz](https://github.com/freezboltz)
* [extractor/airtv] Add extractor by [HobbyistDev](https://github.com/HobbyistDev)
* [extractor/amazonminitv] Add extractors by [GautamMKGarg](https://github.com/GautamMKGarg), [nyuszika7h](https://github.com/nyuszika7h)
* [extractor/beatbump] Add extractors by [Bobscorn](https://github.com/Bobscorn), [pukkandan](https://github.com/pukkandan)
* [extractor/europarl] Add EuroParlWebstream extractor by [HobbyistDev](https://github.com/HobbyistDev)
* [extractor/kanal2] Add extractor by [bashonly](https://github.com/bashonly), [glensc](https://github.com/glensc), [pukkandan](https://github.com/pukkandan)
* [extractor/kankanews] Add extractor by [synthpop123](https://github.com/synthpop123)
* [extractor/kick] Add extractor by [bashonly](https://github.com/bashonly)
* [extractor/mediastream] Add extractor by [HobbyistDev](https://github.com/HobbyistDev), [elyse0](https://github.com/elyse0)
* [extractor/noice] Add NoicePodcast extractor by [HobbyistDev](https://github.com/HobbyistDev)
* [extractor/oneplace] Add OnePlacePodcast extractor by [HobbyistDev](https://github.com/HobbyistDev)
* [extractor/rumble] Add RumbleIE extractor by [flashdagger](https://github.com/flashdagger)
* [extractor/screencastify] Add extractor by [bashonly](https://github.com/bashonly)
* [extractor/trtcocuk] Add extractor by [HobbyistDev](https://github.com/HobbyistDev)
* [extractor/Veoh] Add user extractor by [tntmod54321](https://github.com/tntmod54321)
* [extractor/videoken] Add extractors by [bashonly](https://github.com/bashonly)
* [extractor/webcamerapl] Add extractor by [milkknife](https://github.com/milkknife)
* [extractor/amazon] Add `AmazonReviews` extractor by [bashonly](https://github.com/bashonly)
* [extractor/netverse] Add `NetverseSearch` extractor by [HobbyistDev](https://github.com/HobbyistDev)
* [extractor/vimeo] Add `VimeoProIE` by [bashonly](https://github.com/bashonly), [pukkandan](https://github.com/pukkandan)
* [extractor/xiami] Remove extractors by [synthpop123](https://github.com/synthpop123)
* [extractor/youtube] Add `piped.video` by [Bnyro](https://github.com/Bnyro)
* [extractor/youtube] Consider language in format de-duplication
* [extractor/youtube] Extract DRC formats
* [extractor/youtube] Fix `ytuser:`
* [extractor/youtube] Fix bug in handling of music URLs
* [extractor/youtube] Subtitles cannot be translated to `und`
* [extractor/youtube:tab] Extract metadata from channel items by [coletdjnz](https://github.com/coletdjnz)
* [extractor/ARD] Add vtt subtitles by [CapacitorSet](https://github.com/CapacitorSet)
* [extractor/ArteTV] Extract chapters by [bashonly](https://github.com/bashonly), [iw0nderhow](https://github.com/iw0nderhow)
* [extractor/bandcamp] Add `album_artist` by [stelcodes](https://github.com/stelcodes)
* [extractor/bilibili] Fix `--no-playlist` for anthology
* [extractor/bilibili] Improve `_VALID_URL` by [skbeh](https://github.com/skbeh)
* [extractor/biliintl:series] Make partial download of series faster
* [extractor/BiliLive] Fix extractor
* [extractor/brightcove] Add `BrightcoveNewBaseIE` and fix embed extraction
* [extractor/cda] Support premium and misc improvements by [selfisekai](https://github.com/selfisekai)
* [extractor/ciscowebex] Support password-protected videos by [damianoamatruda](https://github.com/damianoamatruda)
* [extractor/curiositystream] Fix auth by [mnn](https://github.com/mnn)
* [extractor/embedly] Handle vimeo embeds
* [extractor/fifa] Fix Preplay extraction by [dirkf](https://github.com/dirkf)
* [extractor/foxsports] Fix extractor by [bashonly](https://github.com/bashonly)
* [extractor/gronkh] Fix `_VALID_URL` by [muddi900](https://github.com/muddi900)
* [extractor/hotstar] Improve format metadata
* [extractor/iqiyi] Fix `Iq` JS regex by [bashonly](https://github.com/bashonly)
* [extractor/la7] Improve extractor by [nixxo](https://github.com/nixxo)
* [extractor/mediaset] Better embed detection and error messages by [nixxo](https://github.com/nixxo)
* [extractor/mixch] Support `--wait-for-video`
* [extractor/naver] Improve `_VALID_URL` for `NaverNowIE` by [bashonly](https://github.com/bashonly)
* [extractor/naver] Treat fan subtitles as separate language
* [extractor/netverse] Extract comments by [HobbyistDev](https://github.com/HobbyistDev)
* [extractor/nosnl] Add support for /video by [HobbyistDev](https://github.com/HobbyistDev)
* [extractor/odnoklassniki] Extract subtitles by [bashonly](https://github.com/bashonly)
* [extractor/pinterest] Fix extractor by [bashonly](https://github.com/bashonly)
* [extractor/plutotv] Fix videos with non-zero start by [digitall](https://github.com/digitall)
* [extractor/polskieradio] Adapt to next.js redesigns by [selfisekai](https://github.com/selfisekai)
* [extractor/reddit] Add vcodec to fallback format by [chengzhicn](https://github.com/chengzhicn)
* [extractor/reddit] Extract crossposted media by [bashonly](https://github.com/bashonly)
* [extractor/reddit] Extract video embeds in text posts by [bashonly](https://github.com/bashonly)
* [extractor/rutube] Support private videos by [mexus](https://github.com/mexus)
* [extractor/sibnet] Separate from VKIE
* [extractor/slideslive] Fix extractor by [Grub4K](https://github.com/Grub4K), [bashonly](https://github.com/bashonly)
* [extractor/slideslive] Support embeds and slides by [Grub4K](https://github.com/Grub4K), [bashonly](https://github.com/bashonly), [pukkandan](https://github.com/pukkandan)
* [extractor/soundcloud] Support user permalink by [nosoop](https://github.com/nosoop)
* [extractor/spankbang] Fix extractor by [JChris246](https://github.com/JChris246)
* [extractor/stv] Detect DRM
* [extractor/swearnet] Fix description bug
* [extractor/tencent] Fix geo-restricted video by [elyse0](https://github.com/elyse0)
* [extractor/tiktok] Fix subs, `DouyinIE`, improve `_VALID_URL` by [bashonly](https://github.com/bashonly)
* [extractor/tiktok] Update `_VALID_URL`, add `api_hostname` arg by [bashonly](https://github.com/bashonly)
* [extractor/tiktok] Update API hostname by [redraskal](https://github.com/redraskal)
* [extractor/twitcasting] Fix videos with password by [Spicadox](https://github.com/Spicadox), [bashonly](https://github.com/bashonly)
* [extractor/twitter] Heed `--no-playlist` for multi-video tweets by [Grub4K](https://github.com/Grub4K), [bashonly](https://github.com/bashonly)
* [extractor/twitter] Refresh guest token when expired by [Grub4K](https://github.com/Grub4K), [bashonly](https://github.com/bashonly)
* [extractor/twitter:spaces] Add `Referer` to m3u8 by [nixxo](https://github.com/nixxo)
* [extractor/udemy] Fix lectures that have no URL and detect DRM
* [extractor/unsupported] Add more URLs
* [extractor/urplay] Support for audio-only formats by [barsnick](https://github.com/barsnick)
* [extractor/wistia] Improve extension detection by [Grub4K](https://github.com/Grub4K), [bashonly](https://github.com/bashonly), [pukkandan](https://github.com/pukkandan)
* [extractor/yle_areena] Support restricted videos by [docbender](https://github.com/docbender)
* [extractor/youku] Fix extractor by [KurtBestor](https://github.com/KurtBestor)
* [extractor/youporn] Fix metadata by [marieell](https://github.com/marieell)
* [extractor/redgifs] Fix bug in [8c188d5](https://github.com/yt-dlp/yt-dlp/commit/8c188d5d09177ed213a05c900d3523867c5897fd)


### 2022.11.11

* Merge youtube-dl: Upto [commit/de39d12](https://github.com/ytdl-org/youtube-dl/commit/de39d128)
* Backport SSL configuration from Python 3.10 by [coletdjnz](https://github.com/coletdjnz)
* Do more processing in `--flat-playlist`
* Fix `--list` options not implying `-s` in some cases by [Grub4K](https://github.com/Grub4K), [bashonly](https://github.com/bashonly)
* Fix end time of clips by [cruel-efficiency](https://github.com/cruel-efficiency)
* Fix for `formats=None`
* Write API params in debug head
* [outtmpl] Ensure ASCII in json and add option for Unicode
* [SponsorBlock] Add `type` field, obey `--retry-sleep extractor`, relax duration check for large segments
* [SponsorBlock] **Support `chapter` category** by [ajayyy](https://github.com/ajayyy), [pukkandan](https://github.com/pukkandan)
* [ThumbnailsConvertor] Fix filename escaping by [dirkf](https://github.com/dirkf), [pukkandan](https://github.com/pukkandan)
* [ModifyChapters] Handle the entire video being marked for removal
* [embedthumbnail] Fix thumbnail name in mp3 by [How-Bout-No](https://github.com/How-Bout-No)
* [downloader/fragment] HLS download can continue without first fragment
* [cookies] Improve `LenientSimpleCookie` by [Grub4K](https://github.com/Grub4K)
* [jsinterp] Improve separating regex
* [extractor/common] Fix `fatal=False` for `_search_nuxt_data`
* [extractor/common] Improve `_generic_title`
* [extractor/common] Fix `json_ld` type checks by [Grub4K](https://github.com/Grub4K)
* [extractor/generic] Separate embed extraction into own function
* [extractor/generic:quoted-html] Add extractor by [coletdjnz](https://github.com/coletdjnz), [pukkandan](https://github.com/pukkandan)
* [extractor/unsupported] Raise error on known DRM-only sites by [coletdjnz](https://github.com/coletdjnz)
* [utils] `js_to_json`: Improve escape handling by [Grub4K](https://github.com/Grub4K)
* [utils] `strftime_or_none`: Workaround Python bug on Windows
* [utils] `traverse_obj`: Always return list when branching, allow `re.Match` objects by [Grub4K](https://github.com/Grub4K)
* [build, test] Harden workflows' security by [sashashura](https://github.com/sashashura)
* [build] `py2exe`: Migrate to freeze API by [SG5](https://github.com/SG5), [pukkandan](https://github.com/pukkandan)
* [build] Create `armv7l` and `aarch64` releases by [MrOctopus](https://github.com/MrOctopus), [pukkandan](https://github.com/pukkandan)
* [build] Make linux binary truly standalone using `conda` by [mlampe](https://github.com/mlampe)
* [build] Replace `set-output` with `GITHUB_OUTPUT` by [Lesmiscore](https://github.com/Lesmiscore)
* [update] Use error code `100` for update errors
* [compat] Fix `shutils.move` in restricted ACL mode on BSD by [ClosedPort22](https://github.com/ClosedPort22), [pukkandan](https://github.com/pukkandan)
* [docs, devscripts] Document `pyinst`'s argument passthrough by [jahway603](https://github.com/jahway603)
* [test] Allow `extract_flat` in download tests by [coletdjnz](https://github.com/coletdjnz), [pukkandan](https://github.com/pukkandan)
* [cleanup] Misc fixes and cleanup by [pukkandan](https://github.com/pukkandan), [Alienmaster](https://github.com/Alienmaster)
* [extractor/aeon] Add extractor by [DoubleCouponDay](https://github.com/DoubleCouponDay)
* [extractor/agora] Add extractors by [selfisekai](https://github.com/selfisekai)
* [extractor/camsoda] Add extractor by [zulaport](https://github.com/zulaport)
* [extractor/cinetecamilano] Add extractor by [timendum](https://github.com/timendum)
* [extractor/deuxm] Add extractors by [CrankDatSouljaBoy](https://github.com/CrankDatSouljaBoy)
* [extractor/genius] Add extractors by [bashonly](https://github.com/bashonly)
* [extractor/japandiet] Add extractors by [Lesmiscore](https://github.com/Lesmiscore)
* [extractor/listennotes] Add extractor by [lksj](https://github.com/lksj), [pukkandan](https://github.com/pukkandan)
* [extractor/nos.nl] Add extractor by [HobbyistDev](https://github.com/HobbyistDev)
* [extractor/oftv] Add extractors by [DoubleCouponDay](https://github.com/DoubleCouponDay)
* [extractor/podbayfm] Add extractor by [schnusch](https://github.com/schnusch)
* [extractor/qingting] Add extractor by [bashonly](https://github.com/bashonly), [changren-wcr](https://github.com/changren-wcr)
* [extractor/screen9] Add extractor by [tpikonen](https://github.com/tpikonen)
* [extractor/swearnet] Add extractor by [HobbyistDev](https://github.com/HobbyistDev)
* [extractor/YleAreena] Add extractor by [pukkandan](https://github.com/pukkandan), [vitkhab](https://github.com/vitkhab)
* [extractor/zeenews] Add extractor by [m4tu4g](https://github.com/m4tu4g), [pukkandan](https://github.com/pukkandan)
* [extractor/youtube:tab] **Update tab handling for redesign** by [coletdjnz](https://github.com/coletdjnz), [pukkandan](https://github.com/pukkandan)
    * Channel URLs download all uploads of the channel as multiple playlists, separated by tab
* [extractor/youtube] Differentiate between no comments and disabled comments by [coletdjnz](https://github.com/coletdjnz)
* [extractor/youtube] Extract `concurrent_view_count` for livestreams by [coletdjnz](https://github.com/coletdjnz)
* [extractor/youtube] Fix `duration` for premieres by [nosoop](https://github.com/nosoop)
* [extractor/youtube] Fix `live_status` by [coletdjnz](https://github.com/coletdjnz), [pukkandan](https://github.com/pukkandan)
* [extractor/youtube] Ignore incomplete data error for comment replies by [coletdjnz](https://github.com/coletdjnz)
* [extractor/youtube] Improve chapter parsing from description
* [extractor/youtube] Mark videos as fully watched by [bsun0000](https://github.com/bsun0000)
* [extractor/youtube] Update piped instances by [Generator](https://github.com/Generator)
* [extractor/youtube] Update playlist metadata extraction for new layout by [coletdjnz](https://github.com/coletdjnz)
* [extractor/youtube:tab] Fix video metadata from tabs by [coletdjnz](https://github.com/coletdjnz)
* [extractor/youtube:tab] Let `approximate_date` return timestamp
* [extractor/americastestkitchen] Fix extractor by [bashonly](https://github.com/bashonly)
* [extractor/bbc] Support onion domains by [DoubleCouponDay](https://github.com/DoubleCouponDay)
* [extractor/bilibili] Add chapters and misc cleanup by [lockmatrix](https://github.com/lockmatrix), [pukkandan](https://github.com/pukkandan)
* [extractor/bilibili] Fix BilibiliIE and Bangumi extractors by [lockmatrix](https://github.com/lockmatrix), [pukkandan](https://github.com/pukkandan)
* [extractor/bitchute] Better error for geo-restricted videos by [flashdagger](https://github.com/flashdagger)
* [extractor/bitchute] Improve `BitChuteChannelIE` by [flashdagger](https://github.com/flashdagger), [pukkandan](https://github.com/pukkandan)
* [extractor/bitchute] Simplify extractor by [flashdagger](https://github.com/flashdagger), [pukkandan](https://github.com/pukkandan)
* [extractor/cda] Support login through API by [selfisekai](https://github.com/selfisekai)
* [extractor/crunchyroll] Beta is now the only layout by [tejing1](https://github.com/tejing1)
* [extractor/detik] Avoid unnecessary extraction
* [extractor/doodstream] Remove extractor
* [extractor/dplay] Add MotorTrendOnDemand extractor by [bashonly](https://github.com/bashonly)
* [extractor/epoch] Support videos without data-trailer by [gibson042](https://github.com/gibson042), [pukkandan](https://github.com/pukkandan)
* [extractor/fox] Extract thumbnail by [vitkhab](https://github.com/vitkhab)
* [extractor/foxnews] Add `FoxNewsVideo` extractor
* [extractor/hotstar] Add season support by [m4tu4g](https://github.com/m4tu4g)
* [extractor/hotstar] Refactor v1 API calls
* [extractor/iprima] Make json+ld non-fatal by [bashonly](https://github.com/bashonly)
* [extractor/iq] Increase phantomjs timeout
* [extractor/kaltura] Support playlists by [jwoglom](https://github.com/jwoglom), [pukkandan](https://github.com/pukkandan)
* [extractor/lbry] Authenticate with cookies by [flashdagger](https://github.com/flashdagger)
* [extractor/livestreamfails] Support posts by [invertico](https://github.com/invertico)
* [extractor/mlb] Add `MLBArticle` extractor by [HobbyistDev](https://github.com/HobbyistDev)
* [extractor/mxplayer] Improve extractor by [m4tu4g](https://github.com/m4tu4g)
* [extractor/niconico] Always use HTTPS for requests
* [extractor/nzherald] Support new video embed by [coletdjnz](https://github.com/coletdjnz)
* [extractor/odnoklassniki] Support boosty.to embeds by [Lesmiscore](https://github.com/Lesmiscore), [megapro17](https://github.com/megapro17), [pukkandan](https://github.com/pukkandan)
* [extractor/paramountplus] Update API token by [bashonly](https://github.com/bashonly)
* [extractor/reddit] Add fallback format by [bashonly](https://github.com/bashonly)
* [extractor/redgifs] Fix extractors by [bashonly](https://github.com/bashonly), [pukkandan](https://github.com/pukkandan)
* [extractor/redgifs] Refresh auth token for 401 by [endotronic](https://github.com/endotronic), [pukkandan](https://github.com/pukkandan)
* [extractor/rumble] Add HLS formats and extract more metadata by [flashdagger](https://github.com/flashdagger)
* [extractor/sbs] Improve `_VALID_URL` by [bashonly](https://github.com/bashonly)
* [extractor/skyit] Fix extractors by [nixxo](https://github.com/nixxo)
* [extractor/stripchat] Fix hostname for HLS stream by [zulaport](https://github.com/zulaport)
* [extractor/stripchat] Improve error message by [freezboltz](https://github.com/freezboltz)
* [extractor/telegram] Add playlist support and more metadata by [bashonly](https://github.com/bashonly), [bsun0000](https://github.com/bsun0000)
* [extractor/Tnaflix] Fix for HTTP 500 by [SG5](https://github.com/SG5), [pukkandan](https://github.com/pukkandan)
* [extractor/tubitv] Better DRM detection by [bashonly](https://github.com/bashonly)
* [extractor/tvp] Update extractors by [selfisekai](https://github.com/selfisekai)
* [extractor/twitcasting] Fix `data-movie-playlist` extraction by [Lesmiscore](https://github.com/Lesmiscore)
* [extractor/twitter] Add onion site to `_VALID_URL` by [DoubleCouponDay](https://github.com/DoubleCouponDay)
* [extractor/twitter] Add Spaces extractor and GraphQL API by [Grub4K](https://github.com/Grub4K), [bashonly](https://github.com/bashonly), [nixxo](https://github.com/nixxo), [pukkandan](https://github.com/pukkandan)
* [extractor/twitter] Support multi-video posts by [Grub4K](https://github.com/Grub4K)
* [extractor/uktvplay] Fix `_VALID_URL`
* [extractor/viu] Support subtitles of on-screen text by [tkgmomosheep](https://github.com/tkgmomosheep)
* [extractor/VK] Fix playlist URLs by [the-marenga](https://github.com/the-marenga)
* [extractor/vlive] Extract `release_timestamp`
* [extractor/voot] Improve `_VALID_URL` by [freezboltz](https://github.com/freezboltz)
* [extractor/wordpress:mb.miniAudioPlayer] Add embed extractor by [coletdjnz](https://github.com/coletdjnz)
* [extractor/YoutubeWebArchive] Improve metadata extraction by [coletdjnz](https://github.com/coletdjnz)
* [extractor/zee5] Improve `_VALID_URL` by [m4tu4g](https://github.com/m4tu4g)
* [extractor/zenyandex] Fix extractors by [lksj](https://github.com/lksj), [puc9](https://github.com/puc9), [pukkandan](https://github.com/pukkandan)


### 2022.10.04

* Allow a `set` to be passed as `download_archive` by [pukkandan](https://github.com/pukkandan), [bashonly](https://github.com/bashonly)
* Allow open ranges for time ranges by [Lesmiscore](https://github.com/Lesmiscore)
* Allow plugin extractors to replace the built-in ones
* Don't download entire video when no matching `--download-sections`
* Fix `--config-location -`
* Improve [5736d79](https://github.com/yt-dlp/yt-dlp/pull/5044/commits/5736d79172c47ff84740d5720467370a560febad)
* Fix for when playlists don't have `webpage_url`
* Support environment variables in `--ffmpeg-location`
* Workaround `libc_ver` not be available on Windows Store version of Python
* [outtmpl] Curly braces to filter keys by [pukkandan](https://github.com/pukkandan)
* [outtmpl] Make `%s` work in strfformat for all systems
* [jsinterp] Workaround operator associativity issue
* [cookies] Let `_get_mac_keyring_password` fail gracefully
* [cookies] Parse cookies leniently by [Grub4K](https://github.com/Grub4K)
* [phantomjs] Fix bug in [587021c](https://github.com/yt-dlp/yt-dlp/commit/587021cd9f717181b44e881941aca3f8d753758b) by [elyse0](https://github.com/elyse0)
* [downloader/aria2c] Fix filename containing leading whitespace by [std-move](https://github.com/std-move)
* [downloader/ism] Support ec-3 codec by [nixxo](https://github.com/nixxo)
* [extractor] Fix `fatal=False` in `RetryManager`
* [extractor] Improve json-ld extraction
* [extractor] Make `_search_json` able to parse lists
* [extractor] Escape `%` in `representation_id` of m3u8
* [extractor/generic] Pass through referer from json-ld
* [utils] `base_url`: URL paths can contain `&` by [elyse0](https://github.com/elyse0)
* [utils] `js_to_json`: Improve
* [utils] `Popen.run`: Fix default return in binary mode
* [utils] `traverse_obj`: Rewrite, document and add tests by [Grub4K](https://github.com/Grub4K)
* [devscripts] `make_lazy_extractors`: Fix for Docker by [josanabr](https://github.com/josanabr)
* [docs] Misc Improvements
* [cleanup] Misc fixes and cleanup by [pukkandan](https://github.com/pukkandan), [gamer191](https://github.com/gamer191)
* [extractor/24tv.ua] Add extractors by [coletdjnz](https://github.com/coletdjnz)
* [extractor/BerufeTV] Add extractor by [Fabi019](https://github.com/Fabi019)
* [extractor/booyah] Add extractor by [HobbyistDev](https://github.com/HobbyistDev), [elyse0](https://github.com/elyse0)
* [extractor/bundesliga] Add extractor by [Fabi019](https://github.com/Fabi019)
* [extractor/GoPlay] Add extractor by [CNugteren](https://github.com/CNugteren), [basrieter](https://github.com/basrieter), [jeroenj](https://github.com/jeroenj)
* [extractor/iltalehti] Add extractor by [tpikonen](https://github.com/tpikonen)
* [extractor/IsraelNationalNews] Add extractor by [Bobscorn](https://github.com/Bobscorn)
* [extractor/mediaworksnzvod] Add extractor by [coletdjnz](https://github.com/coletdjnz)
* [extractor/MicrosoftEmbed] Add extractor by [DoubleCouponDay](https://github.com/DoubleCouponDay)
* [extractor/nbc] Add NBCStations extractor by [bashonly](https://github.com/bashonly)
* [extractor/onenewsnz] Add extractor by [coletdjnz](https://github.com/coletdjnz)
* [extractor/prankcast] Add extractor by [HobbyistDev](https://github.com/HobbyistDev), [columndeeply](https://github.com/columndeeply)
* [extractor/Smotrim] Add extractor by [Lesmiscore](https://github.com/Lesmiscore), [nikita-moor](https://github.com/nikita-moor)
* [extractor/tencent] Add Iflix extractor by [elyse0](https://github.com/elyse0)
* [extractor/unscripted] Add extractor by [HobbyistDev](https://github.com/HobbyistDev)
* [extractor/adobepass] Add MSO AlticeOne (Optimum TV) by [CplPwnies](https://github.com/CplPwnies)
* [extractor/youtube] **Download `post_live` videos from start** by [Lesmiscore](https://github.com/Lesmiscore), [pukkandan](https://github.com/pukkandan)
* [extractor/youtube] Add support for Shorts audio pivot feed by [coletdjnz](https://github.com/coletdjnz), [pukkandan](https://github.com/pukkandan)
* [extractor/youtube] Detect `lazy-load-for-videos` embeds
* [extractor/youtube] Do not warn on duplicate chapters
* [extractor/youtube] Fix video like count extraction by [coletdjnz](https://github.com/coletdjnz)
* [extractor/youtube] Support changing extraction language by [coletdjnz](https://github.com/coletdjnz)
* [extractor/youtube:tab] Improve continuation items extraction
* [extractor/youtube:tab] Support `reporthistory` page
* [extractor/amazonstore] Fix JSON extraction by [coletdjnz](https://github.com/coletdjnz), [pukkandan](https://github.com/pukkandan)
* [extractor/amazonstore] Retry to avoid captcha page by [Lesmiscore](https://github.com/Lesmiscore)
* [extractor/animeondemand] Remove extractor by [TokyoBlackHole](https://github.com/TokyoBlackHole)
* [extractor/anvato] Fix extractor and refactor by [bashonly](https://github.com/bashonly)
* [extractor/artetv] Remove duplicate stream urls by [Grub4K](https://github.com/Grub4K)
* [extractor/audioboom] Support direct URLs and refactor by [pukkandan](https://github.com/pukkandan), [tpikonen](https://github.com/tpikonen)
* [extractor/bandcamp] Extract `uploader_url`
* [extractor/bilibili] Add space.bilibili extractors by [lockmatrix](https://github.com/lockmatrix)
* [extractor/BilibiliSpace] Fix extractor and better error message by [lockmatrix](https://github.com/lockmatrix)
* [extractor/BiliIntl] Support uppercase lang in `_VALID_URL` by [coletdjnz](https://github.com/coletdjnz)
* [extractor/BiliIntlSeries] Fix `_VALID_URL`
* [extractor/bongacams] Update `_VALID_URL` by [0xGodspeed](https://github.com/0xGodspeed)
* [extractor/crunchyroll:beta] Improve handling of hardsubs by [Grub4K](https://github.com/Grub4K)
* [extractor/detik] Generalize extractors by [HobbyistDev](https://github.com/HobbyistDev), [coletdjnz](https://github.com/coletdjnz)
* [extractor/dplay:italy] Add default authentication by [Timendum](https://github.com/Timendum)
* [extractor/heise] Fix extractor by [coletdjnz](https://github.com/coletdjnz)
* [extractor/holodex] Fix `_VALID_URL` by [LiviaMedeiros](https://github.com/LiviaMedeiros)
* [extractor/hrfensehen] Fix extractor by [snapdgn](https://github.com/snapdgn)
* [extractor/hungama] Add subtitle by [GautamMKGarg](https://github.com/GautamMKGarg), [pukkandan](https://github.com/pukkandan)
* [extractor/instagram] Extract more metadata by [pritam20ps05](https://github.com/pritam20ps05)
* [extractor/JWPlatform] Fix extractor by [coletdjnz](https://github.com/coletdjnz)
* [extractor/malltv] Fix video_id extraction by [HobbyistDev](https://github.com/HobbyistDev)
* [extractor/MLBTV] Detect live streams
* [extractor/motorsport] Support native embeds
* [extractor/Mxplayer] Fix extractor by [itachi-19](https://github.com/itachi-19)
* [extractor/nebula] Add nebula.tv by [tannertechnology](https://github.com/tannertechnology)
* [extractor/nfl] Fix extractor by [bashonly](https://github.com/bashonly)
* [extractor/ondemandkorea] Update `jw_config` regex by [julien-hadleyjack](https://github.com/julien-hadleyjack)
* [extractor/paramountplus] Better DRM detection by [bashonly](https://github.com/bashonly)
* [extractor/patreon] Sort formats
* [extractor/rcs] Fix embed extraction by [coletdjnz](https://github.com/coletdjnz)
* [extractor/redgifs] Fix extractor by [jhwgh1968](https://github.com/jhwgh1968)
* [extractor/rutube] Fix `_EMBED_REGEX` by [coletdjnz](https://github.com/coletdjnz)
* [extractor/RUTV] Fix warnings for livestreams by [Lesmiscore](https://github.com/Lesmiscore)
* [extractor/soundcloud:search] More metadata in `--flat-playlist` by [SuperSonicHub1](https://github.com/SuperSonicHub1)
* [extractor/telegraaf] Use mobile GraphQL API endpoint by [coletdjnz](https://github.com/coletdjnz)
* [extractor/tennistv] Fix timestamp by [zenerdi0de](https://github.com/zenerdi0de)
* [extractor/tiktok] Fix TikTokIE by [bashonly](https://github.com/bashonly)
* [extractor/triller] Fix auth token by [bashonly](https://github.com/bashonly)
* [extractor/trovo] Fix extractors by [Mehavoid](https://github.com/Mehavoid)
* [extractor/tv2] Support new url format by [tobi1805](https://github.com/tobi1805)
* [extractor/web.archive:youtube] Fix `_YT_INITIAL_PLAYER_RESPONSE_RE`
* [extractor/wistia] Add support for channels by [coletdjnz](https://github.com/coletdjnz)
* [extractor/wistia] Match IDs in embed URLs by [bashonly](https://github.com/bashonly)
* [extractor/wordpress:playlist] Add generic embed extractor by [coletdjnz](https://github.com/coletdjnz)
* [extractor/yandexvideopreview] Update `_VALID_URL` by [Grub4K](https://github.com/Grub4K)
* [extractor/zee5] Fix `_VALID_URL` by [m4tu4g](https://github.com/m4tu4g)
* [extractor/zee5] Generate device ids by [freezboltz](https://github.com/freezboltz)


### 2022.09.01

* Add option `--use-extractors`
* Merge youtube-dl: Upto [commit/ed5c44e](https://github.com/ytdl-org/youtube-dl/commit/ed5c44e7)
* Add yt-dlp version to infojson
* Fix `--break-per-url --max-downloads`
* Fix bug in `--alias`
* [cookies] Support firefox container in `--cookies-from-browser` by [bashonly](https://github.com/bashonly), [coletdjnz](https://github.com/coletdjnz), [pukkandan](https://github.com/pukkandan)
* [downloader/external] Smarter detection of executable
* [extractor/generic] Don't return JW player without formats
* [FormatSort] Fix `aext` for `--prefer-free-formats`
* [jsinterp] Various improvements by [pukkandan](https://github.com/pukkandan), [dirkf](https://github.com/dirkf), [elyse0](https://github.com/elyse0)
* [cache] Mechanism to invalidate old cache
* [utils] Add `deprecation_warning`
* [utils] Add `orderedSet_from_options`
* [utils] `Popen`: Restore `LD_LIBRARY_PATH` when using PyInstaller by [Lesmiscore](https://github.com/Lesmiscore)
* [build] `make tar` should not follow `DESTDIR` by [satan1st](https://github.com/satan1st)
* [build] Update pyinstaller by [shirt-dev](https://github.com/shirt-dev)
* [test] Fix `test_youtube_signature`
* [cleanup] Misc fixes and cleanup by [DavidH-2022](https://github.com/DavidH-2022), [MrRawes](https://github.com/MrRawes), [pukkandan](https://github.com/pukkandan)
* [extractor/epoch] Add extractor by [tejasa97](https://github.com/tejasa97)
* [extractor/eurosport] Add extractor by [HobbyistDev](https://github.com/HobbyistDev)
* [extractor/IslamChannel] Add extractors by [Lesmiscore](https://github.com/Lesmiscore)
* [extractor/newspicks] Add extractor by [Lesmiscore](https://github.com/Lesmiscore)
* [extractor/triller] Add extractor by [bashonly](https://github.com/bashonly)
* [extractor/VQQ] Add extractors by [elyse0](https://github.com/elyse0)
* [extractor/youtube] Improvements to nsig extraction
* [extractor/youtube] Fix bug in format sorting
* [extractor/youtube] Update iOS Innertube clients by [SamantazFox](https://github.com/SamantazFox)
* [extractor/youtube] Use device-specific user agent by [coletdjnz](https://github.com/coletdjnz)
* [extractor/youtube] Add `--compat-option no-youtube-prefer-utc-upload-date` by [coletdjnz](https://github.com/coletdjnz)
* [extractor/arte] Bug fix by [cgrigis](https://github.com/cgrigis)
* [extractor/bilibili] Extract `flac` with premium account by [jackyyf](https://github.com/jackyyf)
* [extractor/BiliBiliSearch] Don't sort by date
* [extractor/BiliBiliSearch] Fix infinite loop
* [extractor/bitchute] Mark errors as expected
* [extractor/crunchyroll:beta] Use anonymous access by [tejing1](https://github.com/tejing1)
* [extractor/huya] Fix stream extraction by [ohaiibuzzle](https://github.com/ohaiibuzzle)
* [extractor/medaltv] Fix extraction by [xenova](https://github.com/xenova)
* [extractor/mediaset] Fix embed extraction
* [extractor/mixcloud] All formats are audio-only
* [extractor/rtbf] Fix jwt extraction by [elyse0](https://github.com/elyse0)
* [extractor/screencastomatic] Support `--video-password` by [shreyasminocha](https://github.com/shreyasminocha)
* [extractor/stripchat] Don't modify input URL by [dfaker](https://github.com/dfaker)
* [extractor/uktv] Improve `_VALID_URL` by [dirkf](https://github.com/dirkf)
* [extractor/vimeo:user] Fix `_VALID_URL`


### 2022.08.19

* Fix bug in `--download-archive`
* [jsinterp] **Fix for new youtube players** and related improvements by [dirkf](https://github.com/dirkf), [pukkandan](https://github.com/pukkandan)
* [phantomjs] Add function to execute JS without a DOM by [MinePlayersPE](https://github.com/MinePlayersPE), [pukkandan](https://github.com/pukkandan)
* [build] Exclude devscripts from installs by [Lesmiscore](https://github.com/Lesmiscore)
* [cleanup] Misc fixes and cleanup
* [extractor/youtube] **Add fallback to phantomjs** for nsig
* [extractor/youtube] Fix error reporting of "Incomplete data"
* [extractor/youtube] Improve format sorting for IOS formats
* [extractor/youtube] Improve signature caching
* [extractor/instagram] Fix extraction by [bashonly](https://github.com/bashonly), [pritam20ps05](https://github.com/pritam20ps05)
* [extractor/rai] Minor fix by [nixxo](https://github.com/nixxo)
* [extractor/rtbf] Fix stream extractor by [elyse0](https://github.com/elyse0)
* [extractor/SovietsCloset] Fix extractor by [ChillingPepper](https://github.com/ChillingPepper)
* [extractor/zattoo] Fix Zattoo resellers by [goggle](https://github.com/goggle)

### 2022.08.14

* Merge youtube-dl: Upto [commit/d231b56](https://github.com/ytdl-org/youtube-dl/commit/d231b56)
* [jsinterp] Handle **new youtube signature functions**
* [jsinterp] Truncate error messages
* [extractor] Fix format sorting of `channels`
* [ffmpeg] Disable avconv unless `--prefer-avconv`
* [ffmpeg] Smarter detection of ffprobe filename
* [embedthumbnail] Detect `libatomicparsley.so`
* [ThumbnailsConvertor] Fix conversion after `fixup_webp`
* [utils] Fix `get_compatible_ext`
* [build] Fix changelog
* [update] Set executable bit-mask by [pukkandan](https://github.com/pukkandan), [Lesmiscore](https://github.com/Lesmiscore)
* [devscripts] Fix import
* [docs] Consistent use of `e.g.` by [Lesmiscore](https://github.com/Lesmiscore)
* [cleanup] Misc fixes and cleanup
* [extractor/moview] Add extractor by [HobbyistDev](https://github.com/HobbyistDev)
* [extractor/parler] Add extractor by [palewire](https://github.com/palewire)
* [extractor/patreon] Ignore erroneous media attachments by [coletdjnz](https://github.com/coletdjnz)
* [extractor/truth] Add extractor by [palewire](https://github.com/palewire)
* [extractor/aenetworks] Add formats parameter by [jacobtruman](https://github.com/jacobtruman)
* [extractor/crunchyroll] Improve `_VALID_URL`s
* [extractor/doodstream] Add `wf` domain by [aldoridhoni](https://github.com/aldoridhoni)
* [extractor/facebook] Add reel support by [bashonly](https://github.com/bashonly)
* [extractor/MLB] New extractor by [ischmidt20](https://github.com/ischmidt20)
* [extractor/rai] Misc fixes by [nixxo](https://github.com/nixxo)
* [extractor/toggo] Improve `_VALID_URL` by [masta79](https://github.com/masta79)
* [extractor/tubitv] Extract additional formats by [shirt-dev](https://github.com/shirt-dev)
* [extractor/zattoo] Potential fix for resellers


### 2022.08.08

* **Remove Python 3.6 support**
* Determine merge container better by [pukkandan](https://github.com/pukkandan), [selfisekai](https://github.com/selfisekai)
* Framework for embed detection by [coletdjnz](https://github.com/coletdjnz), [pukkandan](https://github.com/pukkandan)
* Merge youtube-dl: Upto [commit/adb5294](https://github.com/ytdl-org/youtube-dl/commit/adb5294)
* `--compat-option no-live-chat` should disable danmaku
* Fix misleading DRM message
* Import ctypes only when necessary
* Minor bugfixes
* Reject entire playlists faster with `--match-filter`
* Remove filtered entries from `-J`
* Standardize retry mechanism
* Validate `--merge-output-format`
* [downloader] Add average speed to final progress line
* [extractor] Add field `audio_channels`
* [extractor] Support multiple archive ids for one video
* [ffmpeg] Set `ffmpeg_location` in a contextvar
* [FFmpegThumbnailsConvertor] Fix conversion from GIF
* [MetadataParser] Don't set `None` when the field didn't match
* [outtmpl] Smarter replacing of unsupported characters
* [outtmpl] Treat empty values as None in filenames
* [utils] sanitize_open: Allow any IO stream as stdout
* [build, devscripts] Add devscript to set a build variant
* [build] Improve build process by [shirt-dev](https://github.com/shirt-dev)
* [build] Update pyinstaller
* [devscripts] Create `utils` and refactor
* [docs] Clarify `best*`
* [docs] Fix bug report issue template
* [docs] Fix capitalization in references by [christoph-heinrich](https://github.com/christoph-heinrich)
* [cleanup, mhtml] Use imghdr
* [cleanup, utils] Consolidate known media extensions
* [cleanup] Misc fixes and cleanup
* [extractor/angel] Add extractor by [AxiosDeminence](https://github.com/AxiosDeminence)
* [extractor/dplay] Add MotorTrend extractor by [Sipherdrakon](https://github.com/Sipherdrakon)
* [extractor/harpodeon] Add extractor by [eren-kemer](https://github.com/eren-kemer)
* [extractor/holodex] Add extractor by [pukkandan](https://github.com/pukkandan), [sqrtNOT](https://github.com/sqrtNOT)
* [extractor/kompas] Add extractor by [HobbyistDev](https://github.com/HobbyistDev)
* [extractor/rai] Add raisudtirol extractor by [nixxo](https://github.com/nixxo)
* [extractor/tempo] Add extractor by [HobbyistDev](https://github.com/HobbyistDev)
* [extractor/youtube] **Fixes for third party client detection** by [coletdjnz](https://github.com/coletdjnz)
* [extractor/youtube] Add `live_status=post_live` by [lazypete365](https://github.com/lazypete365)
* [extractor/youtube] Extract more format info
* [extractor/youtube] Parse translated subtitles only when requested
* [extractor/youtube, extractor/twitch] Allow waiting for channels to become live
* [extractor/youtube, webvtt] Extract auto-subs from livestream VODs by [fstirlitz](https://github.com/fstirlitz), [pukkandan](https://github.com/pukkandan)
* [extractor/AbemaTVTitle] Implement paging by [Lesmiscore](https://github.com/Lesmiscore)
* [extractor/archiveorg] Improve handling of formats by [coletdjnz](https://github.com/coletdjnz), [pukkandan](https://github.com/pukkandan)
* [extractor/arte] Fix title extraction
* [extractor/arte] **Move to v2 API** by [fstirlitz](https://github.com/fstirlitz), [pukkandan](https://github.com/pukkandan)
* [extractor/bbc] Fix news articles by [ajj8](https://github.com/ajj8)
* [extractor/camtasia] Separate into own extractor by [coletdjnz](https://github.com/coletdjnz)
* [extractor/cloudflarestream] Fix video_id padding by [haobinliang](https://github.com/haobinliang)
* [extractor/crunchyroll] Fix conversion of thumbnail from GIF
* [extractor/crunchyroll] Handle missing metadata correctly by [Burve](https://github.com/Burve), [pukkandan](https://github.com/pukkandan)
* [extractor/crunchyroll:beta] Extract timestamp and fix tests by [tejing1](https://github.com/tejing1)
* [extractor/crunchyroll:beta] Use streams API by [tejing1](https://github.com/tejing1)
* [extractor/doodstream] Support more domains by [Galiley](https://github.com/Galiley)
* [extractor/ESPN] Extract duration by [ischmidt20](https://github.com/ischmidt20)
* [extractor/FIFA] Change API endpoint by [Bricio](https://github.com/Bricio), [yashkc2025](https://github.com/yashkc2025)
* [extractor/globo:article] Remove false positives by [Bricio](https://github.com/Bricio)
* [extractor/Go] Extract timestamp by [ischmidt20](https://github.com/ischmidt20)
* [extractor/hidive] Fix cookie login when netrc is also given by [winterbird-code](https://github.com/winterbird-code)
* [extractor/html5] Separate into own extractor by [coletdjnz](https://github.com/coletdjnz), [pukkandan](https://github.com/pukkandan)
* [extractor/ina] Improve extractor by [elyse0](https://github.com/elyse0)
* [extractor/NaverNow] Change endpoint by [ping](https://github.com/ping)
* [extractor/ninegag] Extract uploader by [DjesonPV](https://github.com/DjesonPV)
* [extractor/NovaPlay] Fix extractor by [Bojidarist](https://github.com/Bojidarist)
* [extractor/orf:radio] Rewrite extractors
* [extractor/patreon] Fix and improve extractors by [coletdjnz](https://github.com/coletdjnz), [pukkandan](https://github.com/pukkandan)
* [extractor/rai] Fix RaiNews extraction by [nixxo](https://github.com/nixxo)
* [extractor/redbee] Unify and update extractors by [elyse0](https://github.com/elyse0)
* [extractor/stripchat] Fix _VALID_URL by [freezboltz](https://github.com/freezboltz)
* [extractor/tubi] Exclude playlists from playlist entries by [sqrtNOT](https://github.com/sqrtNOT)
* [extractor/tviplayer] Improve `_VALID_URL` by [HobbyistDev](https://github.com/HobbyistDev)
* [extractor/twitch] Extract chapters for single chapter VODs by [mpeter50](https://github.com/mpeter50)
* [extractor/vgtv] Support tv.vg.no by [sqrtNOT](https://github.com/sqrtNOT)
* [extractor/vidio] Support embed link by [HobbyistDev](https://github.com/HobbyistDev)
* [extractor/vk] Fix extractor by [Mehavoid](https://github.com/Mehavoid)
* [extractor/WASDTV:record] Fix `_VALID_URL`
* [extractor/xfileshare] Add Referer by [Galiley](https://github.com/Galiley)
* [extractor/YahooJapanNews] Fix extractor by [Lesmiscore](https://github.com/Lesmiscore)
* [extractor/yandexmusic] Extract higher quality format
* [extractor/zee5] Update Device ID by [m4tu4g](https://github.com/m4tu4g)


### 2022.07.18

* Allow users to specify encoding in each config files by [Lesmiscore](https://github.com/Lesmiscore)
* Discard infodict from memory if no longer needed
* Do not allow extractors to return `None`
* Do not load system certificates when `certifi` is used
* Fix rounding of integers in format table
* Improve chapter sanitization
* Skip some fixup if remux/recode is needed by [Lesmiscore](https://github.com/Lesmiscore)
* Support `--no-progress` for `--wait-for-video`
* Fix bug in [612f2be](https://github.com/yt-dlp/yt-dlp/commit/612f2be5d3924540158dfbe5f25d841f04cff8c6)
* [outtmpl] Add alternate form `h` for HTML escaping
* [aes] Add multiple padding modes in CBC by [elyse0](https://github.com/elyse0)
* [extractor/common] Passthrough `errnote=False` to parsers
* [extractor/generic] Remove HEAD request
* [http] Ensure the file handle is always closed
* [ModifyChapters] Modify duration in infodict
* [options] Fix aliases to `--config-location`
* [utils] Fix `get_domain`
* [build] Consistent order for lazy extractors by [lamby](https://github.com/lamby)
* [build] Fix architecture suffix of executables by [odo2063](https://github.com/odo2063)
* [build] Improve `setup.py`
* [update] Do not check `_update_spec` when up to date
* [update] Prepare to remove Python 3.6 support
* [compat] Let PyInstaller detect _legacy module
* [devscripts/update-formulae] Do not change dependency section
* [test] Split download tests so they can be more easily run in CI
* [docs] Improve docstring of `download_ranges` by [FirefoxMetzger](https://github.com/FirefoxMetzger)
* [docs] Improve issue templates
* [build] Fix bug in [6d916fe](https://github.com/yt-dlp/yt-dlp/commit/6d916fe709a38e8c4c69b73843acf170b5165931)
* [cleanup, utils] Refactor parse_codecs
* [cleanup] Misc fixes and cleanup
* [extractor/acfun] Add extractors by [lockmatrix](https://github.com/lockmatrix)
* [extractor/Audiodraft] Add extractors by [Ashish0804](https://github.com/Ashish0804), [fstirlitz](https://github.com/fstirlitz)
* [extractor/cellebrite] Add extractor by [HobbyistDev](https://github.com/HobbyistDev)
* [extractor/detik] Add extractor by [HobbyistDev](https://github.com/HobbyistDev)
* [extractor/hytale] Add extractor by [llamasblade](https://github.com/llamasblade), [pukkandan](https://github.com/pukkandan)
* [extractor/liputan6] Add extractor by [HobbyistDev](https://github.com/HobbyistDev)
* [extractor/mocha] Add extractor by [HobbyistDev](https://github.com/HobbyistDev)
* [extractor/rtl.lu] Add extractor by [HobbyistDev](https://github.com/HobbyistDev)
* [extractor/rtvsl] Add extractor by [iw0nderhow](https://github.com/iw0nderhow), [pukkandan](https://github.com/pukkandan)
* [extractor/StarTrek] Add extractor by [scy](https://github.com/scy)
* [extractor/syvdk] Add extractor by [misaelaguayo](https://github.com/misaelaguayo)
* [extractor/theholetv] Add extractor by [dosy4ev](https://github.com/dosy4ev)
* [extractor/TubeTuGraz] Add extractor by [Ferdi265](https://github.com/Ferdi265), [pukkandan](https://github.com/pukkandan)
* [extractor/tviplayer] Add extractor by [HobbyistDev](https://github.com/HobbyistDev)
* [extractor/wetv] Add extractors by [elyse0](https://github.com/elyse0)
* [extractor/wikimedia] Add extractor by [EhtishamSabir](https://github.com/EhtishamSabir), [pukkandan](https://github.com/pukkandan)
* [extractor/youtube] Fix duration check for post-live manifestless mode
* [extractor/youtube] More metadata for storyboards by [ftk](https://github.com/ftk)
* [extractor/bigo] Fix extractor by [Lesmiscore](https://github.com/Lesmiscore)
* [extractor/BiliIntl] Fix subtitle extraction by [MinePlayersPE](https://github.com/MinePlayersPE)
* [extractor/crunchyroll] Improve `_VALID_URL`
* [extractor/fifa] Fix extractor by [ischmidt20](https://github.com/ischmidt20)
* [extractor/instagram] Fix post/story extractors by [pritam20ps05](https://github.com/pritam20ps05), [pukkandan](https://github.com/pukkandan)
* [extractor/iq] Set language correctly for Korean subtitles
* [extractor/MangoTV] Fix subtitle languages
* [extractor/Netverse] Improve playlist extractor by [HobbyistDev](https://github.com/HobbyistDev)
* [extractor/philharmoniedeparis] Fix extractor by [sqrtNOT](https://github.com/sqrtNOT)
* [extractor/Trovo] Fix extractor by [u-spec-png](https://github.com/u-spec-png)
* [extractor/twitch] Support storyboards for VODs by [ftk](https://github.com/ftk)
* [extractor/WatchESPN] Improve `_VALID_URL` by [IONECarter](https://github.com/IONECarter), [dirkf](https://github.com/dirkf)
* [extractor/WSJArticle] Fix video id extraction by [sqrtNOT](https://github.com/sqrtNOT)
* [extractor/Ximalaya] Fix extractors by [lockmatrix](https://github.com/lockmatrix)
* [cleanup, extractor/youtube] Fix tests by [sheerluck](https://github.com/sheerluck)


### 2022.06.29

* Fix `--downloader native`
* Fix `section_end` of clips
* Fix playlist error handling
* Sanitize `chapters`
* [extractor] Fix `_create_request` when headers is None
* [extractor] Fix empty `BaseURL` in MPD
* [ffmpeg] Write full output to debug on error
* [hls] Warn user when trying to download live HLS
* [options] Fix `parse_known_args` for `--`
* [utils] Fix inconsistent default handling between HTTP and HTTPS requests by [coletdjnz](https://github.com/coletdjnz)
* [build] Draft release until complete
* [build] Fix release tag commit
* [build] Standalone x64 builds for MacOS 10.9 by [StefanLobbenmeier](https://github.com/StefanLobbenmeier)
* [update] Ability to set a maximum version for specific variants
* [compat] Fix `compat.WINDOWS_VT_MODE`
* [compat] Remove deprecated functions from core code
* [compat] Remove more functions
* [cleanup, extractor] Reduce direct use of `_downloader`
* [cleanup] Consistent style for file heads
* [cleanup] Fix some typos by [crazymoose77756](https://github.com/crazymoose77756)
* [cleanup] Misc fixes and cleanup
* [extractor/Scrolller] Add extractor by [LunarFang416](https://github.com/LunarFang416)
* [extractor/ViMP] Add playlist extractor by [FestplattenSchnitzel](https://github.com/FestplattenSchnitzel)
* [extractor/fuyin] Add extractor by [HobbyistDev](https://github.com/HobbyistDev)
* [extractor/livestreamfails] Add extractor by [nomevi](https://github.com/nomevi)
* [extractor/premiershiprugby] Add extractor by [HobbyistDev](https://github.com/HobbyistDev)
* [extractor/steam] Add broadcast extractor by [HobbyistDev](https://github.com/HobbyistDev)
* [extractor/youtube] Mark videos as fully watched by [Brett824](https://github.com/Brett824)
* [extractor/CWTV] Extract thumbnail by [ischmidt20](https://github.com/ischmidt20)
* [extractor/ViMP] Add thumbnail and support more sites by [FestplattenSchnitzel](https://github.com/FestplattenSchnitzel)
* [extractor/dropout] Support cookies and login only as needed by [pingiun](https://github.com/pingiun), [pukkandan](https://github.com/pukkandan)
* [extractor/ertflix] Improve `_VALID_URL`
* [extractor/lbry] Use HEAD request for redirect URL by [flashdagger](https://github.com/flashdagger)
* [extractor/mediaset] Improve `_VALID_URL`
* [extractor/npr] Implement [e50c350](https://github.com/yt-dlp/yt-dlp/commit/e50c3500b43d80e4492569c4b4523c4379c6fbb2) differently
* [extractor/tennistv] Rewrite extractor by [pukkandan](https://github.com/pukkandan), [zenerdi0de](https://github.com/zenerdi0de)

### 2022.06.22.1

* [build] Fix updating homebrew formula

### 2022.06.22

* [**Deprecate support for Python 3.6**](https://github.com/yt-dlp/yt-dlp/issues/3764#issuecomment-1154051119)
* **Add option `--download-sections` to download video partially**
    * Chapter regex and time ranges are accepted, e.g. `--download-sections *1:10-2:20`
* Add option `--alias`
* Add option `--lazy-playlist` to process entries as they are received
* Add option `--retry-sleep`
* Add slicing notation to `--playlist-items`
    * Adds support for negative indices and step
    * Add `-I` as alias for `--playlist-index`
    * Makes `--playlist-start`, `--playlist-end`, `--playlist-reverse`, `--no-playlist-reverse` redundant
* `--config-location -` to provide options interactively
* [build] Add Linux standalone builds
* [update] Self-restart after update
* Merge youtube-dl: Upto [commit/8a158a9](https://github.com/ytdl-org/youtube-dl/commit/8a158a9)
* Add `--no-update`
* Allow extractors to specify section_start/end for clips
* Do not print progress to `stderr` with `-q`
* Ensure pre-processor errors do not block video download
* Fix `--simulate --max-downloads`
* Improve error handling of bad config files
* Return an error code if update fails
* Fix bug in [3a408f9](https://github.com/yt-dlp/yt-dlp/commit/3a408f9d199127ca2626359e21a866a09ab236b3)
* [ExtractAudio] Allow conditional conversion
* [ModifyChapters] Fix repeated removal of small segments
* [ThumbnailsConvertor] Allow conditional conversion
* [cookies] Detect profiles for cygwin/BSD by [moench-tegeder](https://github.com/moench-tegeder)
* [dash] Show fragment count with `--live-from-start` by [flashdagger](https://github.com/flashdagger)
* [extractor] Add `_search_json` by [coletdjnz](https://github.com/coletdjnz), [pukkandan](https://github.com/pukkandan)
* [extractor] Add `default` parameter to `_search_json` by [coletdjnz](https://github.com/coletdjnz), [pukkandan](https://github.com/pukkandan)
* [extractor] Add dev option `--load-pages`
* [extractor] Handle `json_ld` with multiple `@type`s
* [extractor] Import `_ALL_CLASSES` lazily
* [extractor] Recognize `src` attribute from HTML5 media elements by [Lesmiscore](https://github.com/Lesmiscore)
* [extractor/generic] Revert e6ae51c123897927eb3c9899923d8ffd31c7f85d
* [f4m] Bugfix
* [ffmpeg] Check version lazily
* [jsinterp] Some optimizations and refactoring by [dirkf](https://github.com/dirkf), [pukkandan](https://github.com/pukkandan)
* [utils] Improve performance using `functools.cache`
* [utils] Send HTTP/1.1 ALPN extension by [coletdjnz](https://github.com/coletdjnz)
* [utils] `ExtractorError`: Fix `exc_info`
* [utils] `ISO3166Utils`: Add `EU` and `AP`
* [utils] `Popen`: Refactor to use contextmanager
* [utils] `locked_file`: Fix for PyPy on Windows
* [update] Expose more functionality to API
* [update] Use `.git` folder to distinguish `source`/`unknown`
* [compat] Add `functools.cached_property`
* [test] Fix `FakeYDL` signatures by [coletdjnz](https://github.com/coletdjnz)
* [docs] Improvements
* [cleanup, ExtractAudio] Refactor
* [cleanup, downloader] Refactor `report_progress`
* [cleanup, extractor] Refactor `_download_...` methods
* [cleanup, extractor] Rename `extractors.py` to `_extractors.py`
* [cleanup, utils] Don't use kwargs for `format_field`
* [cleanup, build] Refactor
* [cleanup, docs] Re-indent "Usage and Options" section
* [cleanup] Deprecate `YoutubeDL.parse_outtmpl`
* [cleanup] Misc fixes and cleanup by [Lesmiscore](https://github.com/Lesmiscore), [MrRawes](https://github.com/MrRawes), [christoph-heinrich](https://github.com/christoph-heinrich), [flashdagger](https://github.com/flashdagger), [gamer191](https://github.com/gamer191), [kwconder](https://github.com/kwconder), [pukkandan](https://github.com/pukkandan)
* [extractor/DailyWire] Add extractors by [HobbyistDev](https://github.com/HobbyistDev), [pukkandan](https://github.com/pukkandan)
* [extractor/fourzerostudio] Add extractors by [Lesmiscore](https://github.com/Lesmiscore)
* [extractor/GoogleDrive] Add folder extractor by [evansp](https://github.com/evansp), [pukkandan](https://github.com/pukkandan)
* [extractor/MirrorCoUK] Add extractor by [LunarFang416](https://github.com/LunarFang416), [pukkandan](https://github.com/pukkandan)
* [extractor/atscaleconfevent] Add extractor by [Ashish0804](https://github.com/Ashish0804)
* [extractor/freetv] Add extractor by [elyse0](https://github.com/elyse0)
* [extractor/ixigua] Add Extractor by [HobbyistDev](https://github.com/HobbyistDev)
* [extractor/kicker.de] Add extractor by [HobbyistDev](https://github.com/HobbyistDev)
* [extractor/netverse] Add extractors by [HobbyistDev](https://github.com/HobbyistDev), [pukkandan](https://github.com/pukkandan)
* [extractor/playsuisse] Add extractor by [pukkandan](https://github.com/pukkandan), [sbor23](https://github.com/sbor23)
* [extractor/substack] Add extractor by [elyse0](https://github.com/elyse0)
* [extractor/youtube] **Support downloading clips**
* [extractor/youtube] Add `innertube_host` and `innertube_key` extractor args by [coletdjnz](https://github.com/coletdjnz)
* [extractor/youtube] Add warning for PostLiveDvr
* [extractor/youtube] Bring back `_extract_chapters_from_description`
* [extractor/youtube] Extract `comment_count` from webpage
* [extractor/youtube] Fix `:ytnotifications` extractor by [coletdjnz](https://github.com/coletdjnz)
* [extractor/youtube] Fix initial player response extraction by [coletdjnz](https://github.com/coletdjnz), [pukkandan](https://github.com/pukkandan)
* [extractor/youtube] Fix live chat for videos with content warning by [coletdjnz](https://github.com/coletdjnz)
* [extractor/youtube] Make signature extraction non-fatal
* [extractor/youtube:tab] Detect `videoRenderer` in `_post_thread_continuation_entries`
* [extractor/BiliIntl] Fix metadata extraction
* [extractor/BiliIntl] Fix subtitle extraction by [HobbyistDev](https://github.com/HobbyistDev)
* [extractor/FranceCulture] Fix extractor by [aurelg](https://github.com/aurelg), [pukkandan](https://github.com/pukkandan)
* [extractor/PokemonSoundLibrary] Remove extractor by [Lesmiscore](https://github.com/Lesmiscore)
* [extractor/StreamCZ] Fix extractor by [adamanldo](https://github.com/adamanldo), [dirkf](https://github.com/dirkf)
* [extractor/WatchESPN] Support free videos and BAM_DTC by [ischmidt20](https://github.com/ischmidt20)
* [extractor/animelab] Remove extractor by [gamer191](https://github.com/gamer191)
* [extractor/bloomberg] Change playback endpoint by [m4tu4g](https://github.com/m4tu4g)
* [extractor/ccc] Extract view_count by [vkorablin](https://github.com/vkorablin)
* [extractor/crunchyroll:beta] Fix extractor after API change by [Burve](https://github.com/Burve), [tejing1](https://github.com/tejing1)
* [extractor/curiositystream] Get `auth_token` from cookie by [mnn](https://github.com/mnn)
* [extractor/digitalconcerthall] Fix extractor by [ZhymabekRoman](https://github.com/ZhymabekRoman)
* [extractor/dropbox] Extract the correct `mountComponent`
* [extractor/dropout] Login is not mandatory
* [extractor/duboku] Fix for hostname change by [mozbugbox](https://github.com/mozbugbox)
* [extractor/espn] Add `WatchESPN` extractor by [ischmidt20](https://github.com/ischmidt20), [pukkandan](https://github.com/pukkandan)
* [extractor/expressen] Fix extractor by [aejdl](https://github.com/aejdl)
* [extractor/foxnews] Update embed extraction by [elyse0](https://github.com/elyse0)
* [extractor/ina] Fix extractor by [elyse0](https://github.com/elyse0)
* [extractor/iwara:user] Make paging better by [Lesmiscore](https://github.com/Lesmiscore)
* [extractor/jwplatform] Look for `data-video-jw-id`
* [extractor/lbry] Update livestream API by [flashdagger](https://github.com/flashdagger)
* [extractor/mediaset] Improve `_VALID_URL`
* [extractor/naver] Add `navernow` extractor by [ping](https://github.com/ping)
* [extractor/niconico:series] Fix extractor by [sqrtNOT](https://github.com/sqrtNOT)
* [extractor/npr] Use stream url from json-ld by [r5d](https://github.com/r5d)
* [extractor/pornhub] Extract `uploader_id` field by [Lesmiscore](https://github.com/Lesmiscore)
* [extractor/radiofrance] Add more radios by [bubbleguuum](https://github.com/bubbleguuum)
* [extractor/rumble] Detect JS embed
* [extractor/rumble] Extract subtitles by [fstirlitz](https://github.com/fstirlitz)
* [extractor/southpark] Add `southpark.lat` extractor by [darkxex](https://github.com/darkxex)
* [extractor/spotify:show] Fix extractor
* [extractor/tiktok] Detect embeds
* [extractor/tiktok] Extract `SIGI_STATE` by [dirkf](https://github.com/dirkf), [pukkandan](https://github.com/pukkandan), [sulyi](https://github.com/sulyi)
* [extractor/tver] Fix extractor by [Lesmiscore](https://github.com/Lesmiscore)
* [extractor/vevo] Fix extractor by [Lesmiscore](https://github.com/Lesmiscore)
* [extractor/yahoo:gyao] Fix extractor
* [extractor/zattoo] Fix live streams by [miseran](https://github.com/miseran)
* [extractor/zdf] Improve format sorting by [elyse0](https://github.com/elyse0)


### 2022.05.18

* Add support for SSL client certificate authentication by [coletdjnz](https://github.com/coletdjnz), [dirkf](https://github.com/dirkf)
    * Adds `--client-certificate`, `--client-certificate-key`, `--client-certificate-password`
* Add `--match-filter -` to interactively ask for each video
* `--max-downloads` should obey `--break-per-input`
* Allow use of weaker ciphers with `--legacy-server-connect`
* Don't imply `-s` for later stages of `-O`
* Fix `--date today`
* Fix `--skip-unavailable-fragments`
* Fix color in `-q -F`
* Fix redirect HTTP method handling by [coletdjnz](https://github.com/coletdjnz)
* Improve `--clean-infojson`
* Remove warning for videos with an empty title
* Run `FFmpegFixupM3u8PP` for live-streams if needed
* Show name of downloader in verbose log
* [cookies] Allow `cookiefile` to be a text stream
* [cookies] Report progress when importing cookies
* [downloader/ffmpeg] Specify headers for each URL by [elyse0](https://github.com/elyse0)
* [fragment] Do not change chunk-size when `--test`
* [fragment] Make single thread download work for `--live-from-start` by [Lesmiscore](https://github.com/Lesmiscore)
* [hls] Fix `byte_range` for `EXT-X-MAP` fragment by [fstirlitz](https://github.com/fstirlitz)
* [http] Fix retrying on read timeout by [coletdjnz](https://github.com/coletdjnz)
* [ffmpeg] Fix features detection
* [EmbedSubtitle] Enable for more video extensions
* [EmbedThumbnail] Disable thumbnail conversion for mkv by [evansp](https://github.com/evansp)
* [EmbedThumbnail] Do not obey `-k`
* [EmbedThumbnail] Do not remove id3v1 tags
* [FFmpegMetadata] Remove `\0` from metadata
* [FFmpegMetadata] Remove filename from attached info-json
* [FixupM3u8] Obey `--hls-prefer-mpegts`
* [Sponsorblock] Don't crash when duration is unknown
* [XAttrMetadata] Refactor and document dependencies
* [extractor] Document netrc machines
* [extractor] Update `manifest_url`s after redirect by [elyse0](https://github.com/elyse0)
* [extractor] Update dash `manifest_url` after redirects by [elyse0](https://github.com/elyse0)
* [extractor] Use `classmethod`/`property` where possible
* [generic] Refactor `_extract_rss`
* [utils] `is_html`: Handle double BOM
* [utils] `locked_file`: Ignore illegal seek on `truncate` by [jakeogh](https://github.com/jakeogh)
* [utils] `sanitize_path`: Fix when path is empty string
* [utils] `write_string`: Workaround newline issue in `conhost`
* [utils] `certifi`: Make sure the pem file exists
* [utils] Fix `WebSocketsWrapper`
* [utils] `locked_file`: Do not give executable bits for newly created files by [Lesmiscore](https://github.com/Lesmiscore)
* [utils] `YoutubeDLCookieJar`: Detect and reject JSON file by [Lesmiscore](https://github.com/Lesmiscore)
* [test] Convert warnings into errors and fix some existing warnings by [fstirlitz](https://github.com/fstirlitz)
* [dependencies] Create module with all dependency imports
* [compat] Split into sub-modules by [fstirlitz](https://github.com/fstirlitz), [pukkandan](https://github.com/pukkandan)
* [compat] Implement `compat.imghdr`
* [build] Add `make uninstall` by [MrRawes](https://github.com/MrRawes)
* [build] Avoid use of `install -D`
* [build] Fix `Makefile` by [putnam](https://github.com/putnam)
* [build] Fix `--onedir` on macOS
* [build] Add more test-runners
* [cleanup] Deprecate some compat vars by [fstirlitz](https://github.com/fstirlitz), [pukkandan](https://github.com/pukkandan)
* [cleanup] Remove unused code paths, extractors, scripts and tests by [fstirlitz](https://github.com/fstirlitz)
* [cleanup] Upgrade syntax (`pyupgrade`) and sort imports (`isort`)
* [cleanup, docs, build] Misc fixes
* [BilibiliLive] Add extractor by [HE7086](https://github.com/HE7086), [pukkandan](https://github.com/pukkandan)
* [Fifa] Add Extractor by [Bricio](https://github.com/Bricio)
* [goodgame] Add extractor by [nevack](https://github.com/nevack)
* [gronkh] Add playlist extractors by [hatienl0i261299](https://github.com/hatienl0i261299)
* [icareus] Add extractor by [tpikonen](https://github.com/tpikonen), [pukkandan](https://github.com/pukkandan)
* [iwara] Add playlist extractors by [i6t](https://github.com/i6t)
* [Likee] Add extractor by [hatienl0i261299](https://github.com/hatienl0i261299)
* [masters] Add extractor by [m4tu4g](https://github.com/m4tu4g)
* [nebula] Add support for subscriptions by [hheimbuerger](https://github.com/hheimbuerger)
* [Podchaser] Add extractors by [connercsbn](https://github.com/connercsbn)
* [rokfin:search] Add extractor by [P-reducible](https://github.com/P-reducible), [pukkandan](https://github.com/pukkandan)
* [youtube] Add `:ytnotifications` extractor by [krichbanana](https://github.com/krichbanana)
* [youtube] Add YoutubeStoriesIE (`ytstories:<channel UCID>`) by [coletdjnz](https://github.com/coletdjnz)
* [ZingMp3] Add chart and user extractors by [hatienl0i261299](https://github.com/hatienl0i261299)
* [adn] Update AES key by [elyse0](https://github.com/elyse0)
* [adobepass] Allow cookies for authenticating MSO
* [bandcamp] Exclude merch links by [Yipten](https://github.com/Yipten)
* [chingari] Fix archiving and tests
* [DRTV] Improve `_VALID_URL` by [vertan](https://github.com/vertan)
* [facebook] Improve thumbnail extraction by [Wikidepia](https://github.com/Wikidepia)
* [fc2] Stop heatbeating once FFmpeg finishes by [Lesmiscore](https://github.com/Lesmiscore)
* [Gofile] Fix extraction and support password-protected links by [mehq](https://github.com/mehq)
* [hotstar, cleanup] Refactor extractors
* [InfoQ] Don't fail on missing audio format by [evansp](https://github.com/evansp)
* [Jamendo] Extract more metadata by [evansp](https://github.com/evansp)
* [kaltura] Update API calls by [flashdagger](https://github.com/flashdagger)
* [KhanAcademy] Fix extractor by [rand-net](https://github.com/rand-net)
* [LCI] Fix extractor by [MarwenDallel](https://github.com/MarwenDallel)
* [lrt] Support livestreams by [GiedriusS](https://github.com/GiedriusS)
* [niconico] Set `expected_protocol` to a public field
* [Niconico] Support 2FA by [ekangmonyet](https://github.com/ekangmonyet)
* [Olympics] Fix format extension
* [openrec:movie] Enable fallback for /movie/ URLs
* [PearVideo] Add fallback for formats by [hatienl0i261299](https://github.com/hatienl0i261299)
* [radiko] Fix extractor by [Lesmiscore](https://github.com/Lesmiscore)
* [rai] Add `release_year`
* [reddit] Prevent infinite loop
* [rokfin] Implement login by [P-reducible](https://github.com/P-reducible), [pukkandan](https://github.com/pukkandan)
* [ruutu] Support hs.fi embeds by [tpikonen](https://github.com/tpikonen), [pukkandan](https://github.com/pukkandan)
* [spotify] Detect iframe embeds by [fstirlitz](https://github.com/fstirlitz)
* [telegram] Fix metadata extraction
* [tmz, cleanup] Update tests by [diegorodriguezv](https://github.com/diegorodriguezv)
* [toggo] Fix `_VALID_URL` by [ca-za](https://github.com/ca-za)
* [trovo] Update to new API by [nyuszika7h](https://github.com/nyuszika7h)
* [TVer] Improve extraction by [Lesmiscore](https://github.com/Lesmiscore)
* [twitcasting] Pass headers for each formats by [Lesmiscore](https://github.com/Lesmiscore)
* [VideocampusSachsen] Improve extractor by [FestplattenSchnitzel](https://github.com/FestplattenSchnitzel)
* [vimeo] Fix extractors
* [wat] Fix extraction of multi-language videos and subtitles by [elyse0](https://github.com/elyse0)
* [wistia] Fix `_VALID_URL` by [dirkf](https://github.com/dirkf)
* [youtube, cleanup] Minor refactoring by [coletdjnz](https://github.com/coletdjnz), [pukkandan](https://github.com/pukkandan)
* [youtube] Added piped instance urls by [JordanWeatherby](https://github.com/JordanWeatherby)
* [youtube] Deprioritize auto-generated thumbnails
* [youtube] Deprioritize format 22 (often damaged)
* [youtube] Fix episode metadata extraction
* [zee5] Fix extractor by [Ashish0804](https://github.com/Ashish0804)
* [zingmp3, cleanup] Refactor extractors


### 2022.04.08

* Use certificates from `certifi` if installed by [coletdjnz](https://github.com/coletdjnz)
* Treat multiple `--match-filters` as OR
* File locking improvements:
    * Do not lock downloading file on Windows
    * Do not prevent download if locking is unsupported
    * Do not truncate files before locking by [jakeogh](https://github.com/jakeogh), [pukkandan](https://github.com/pukkandan)
    * Fix non-blocking non-exclusive lock
* De-prioritize automatic-subtitles when no `--sub-lang` is given
* Exit after `--dump-user-agent`
* Fallback to video-only format when selecting by extension
* Fix `--abort-on-error` for subtitles
* Fix `--no-overwrite` for playlist infojson
* Fix `--print` with `--ignore-no-formats` when url is `None` by [flashdagger](https://github.com/flashdagger)
* Fix `--sleep-interval`
* Fix `--throttled-rate`
* Fix `autonumber`
* Fix case of `http_headers`
* Fix filepath sanitization in `--print-to-file`
* Handle float in `--wait-for-video`
* Ignore `mhtml` formats from `-f mergeall`
* Ignore format-specific fields in initial pass of `--match-filter`
* Protect stdout from unexpected progress and console-title
* Remove `Accept-Encoding` header from `std_headers` by [coletdjnz](https://github.com/coletdjnz)
* Remove incorrect warning for `--dateafter`
* Show warning when all media formats have DRM
* [downloader] Fix invocation of `HttpieFD`
* [http] Fix #3215
* [http] Reject broken range before request by [Lesmiscore](https://github.com/Lesmiscore), [Jules-A](https://github.com/Jules-A), [pukkandan](https://github.com/pukkandan)
* [fragment] Read downloaded fragments only when needed by [Lesmiscore](https://github.com/Lesmiscore)
* [http] Retry on more errors by [coletdjnz](https://github.com/coletdjnz)
* [mhtml] Fix fragments with absolute urls by [coletdjnz](https://github.com/coletdjnz)
* [extractor] Add `_perform_login` function
* [extractor] Allow control characters inside json
* [extractor] Support merging subtitles with data by [coletdjnz](https://github.com/coletdjnz)
* [generic] Extract subtitles from video.js by [Lesmiscore](https://github.com/Lesmiscore)
* [ffmpeg] Cache version data
* [FFmpegConcat] Ensure final directory exists
* [FfmpegMetadata] Write id3v1 tags
* [FFmpegVideoConvertor] Add more formats to `--remux-video`
* [FFmpegVideoConvertor] Ensure all streams are copied
* [MetadataParser] Validate outtmpl early
* [outtmpl] Fix replacement/default when used with alternate
* [outtmpl] Limit changes during sanitization
* [phantomjs] Fix bug
* [test] Add `test_locked_file`
* [utils] `format_decimal_suffix`: Fix for very large numbers by [s0u1h](https://github.com/s0u1h)
* [utils] `traverse_obj`: Allow filtering by value
* [utils] Add `filter_dict`, `get_first`, `try_call`
* [utils] ExtractorError: Fix for older python versions
* [utils] WebSocketsWrapper: Allow omitting `__enter__` invocation by [Lesmiscore](https://github.com/Lesmiscore)
* [docs] Add an `.editorconfig` file by [fstirlitz](https://github.com/fstirlitz)
* [docs] Clarify the exact `BSD` license of dependencies by [MrRawes](https://github.com/MrRawes)
* [docs] Minor improvements by [pukkandan](https://github.com/pukkandan), [cffswb](https://github.com/cffswb), [danielyli](https://github.com/danielyli)
* [docs] Remove readthedocs
* [build] Add `requirements.txt` to pip distributions
* [cleanup, postprocessor] Create `_download_json`
* [cleanup, vimeo] Fix tests
* [cleanup] Misc fixes and minor cleanup
* [cleanup] Use `_html_extract_title`
* [AfreecaTV] Add `AfreecaTVUserIE` by [hatienl0i261299](https://github.com/hatienl0i261299)
* [arte] Add `format_note` to m3u8 formats
* [azmedien] Add TVO Online to supported hosts by [1-Byte](https://github.com/1-Byte)
* [BanBye] Add extractor by [mehq](https://github.com/mehq)
* [bilibili] Fix extraction of title with quotes by [dzek69](https://github.com/dzek69)
* [Craftsy] Add extractor by [Bricio](https://github.com/Bricio)
* [Cybrary] Add extractor by [aaearon](https://github.com/aaearon)
* [Huya] Add extractor by [hatienl0i261299](https://github.com/hatienl0i261299)
* [ITProTV] Add extractor by [aaearon](https://github.com/aaearon)
* [Jable] Add extractors by [mehq](https://github.com/mehq)
* [LastFM] Add extractors by [mehq](https://github.com/mehq)
* [Moviepilot] Add extractor by [panatexxa](https://github.com/panatexxa)
* [panopto] Add extractors by [coletdjnz](https://github.com/coletdjnz), [kmark](https://github.com/kmark)
* [PokemonSoundLibrary] Add extractor by [Lesmiscore](https://github.com/Lesmiscore)
* [WasdTV] Add extractor by [un-def](https://github.com/un-def), [hatienl0i261299](https://github.com/hatienl0i261299)
* [adobepass] Fix Suddenlink MSO by [CplPwnies](https://github.com/CplPwnies)
* [afreecatv] Match new vod url by [wlritchi](https://github.com/wlritchi)
* [AZMedien] Support `tv.telezueri.ch` by [goggle](https://github.com/goggle)
* [BiliIntl] Support user-generated videos by [wlritchi](https://github.com/wlritchi)
* [BRMediathek] Fix VALID_URL
* [crunchyroll:playlist] Implement beta API by [tejing1](https://github.com/tejing1)
* [crunchyroll] Fix inheritance
* [daftsex] Fix extractor by [Soebb](https://github.com/Soebb)
* [dailymotion] Support `geo.dailymotion.com` by [hatienl0i261299](https://github.com/hatienl0i261299)
* [ellentube] Extract subtitles from manifest
* [elonet] Rewrite extractor by [Fam0r](https://github.com/Fam0r), [pukkandan](https://github.com/pukkandan)
* [fptplay] Fix metadata extraction by [hatienl0i261299](https://github.com/hatienl0i261299)
* [FranceCulture] Support playlists by [bohwaz](https://github.com/bohwaz)
* [go, viu] Extract subtitles from the m3u8 manifest by [fstirlitz](https://github.com/fstirlitz)
* [Imdb] Improve extractor by [hatienl0i261299](https://github.com/hatienl0i261299)
* [MangoTV] Improve extractor by [hatienl0i261299](https://github.com/hatienl0i261299)
* [Nebula] Fix bug in 52efa4b31200119adaa8acf33e50b84fcb6948f0
* [niconico] Fix extraction of thumbnails and uploader (#3266)
* [niconico] Rewrite NiconicoIE by [Lesmiscore](https://github.com/Lesmiscore)
* [nitter] Minor fixes and update instance list by [foghawk](https://github.com/foghawk)
* [NRK] Extract timestamp by [hatienl0i261299](https://github.com/hatienl0i261299)
* [openrec] Download archived livestreams by [Lesmiscore](https://github.com/Lesmiscore)
* [openrec] Refactor extractors by [Lesmiscore](https://github.com/Lesmiscore)
* [panopto] Improve subtitle extraction and support slides by [coletdjnz](https://github.com/coletdjnz)
* [ParamountPlus, CBS] Change VALID_URL by [Sipherdrakon](https://github.com/Sipherdrakon)
* [ParamountPlusSeries] Support multiple pages by [dodrian](https://github.com/dodrian)
* [Piapro] Extract description with break lines by [Lesmiscore](https://github.com/Lesmiscore)
* [rai] Fix extraction of http formas by [nixxo](https://github.com/nixxo)
* [rumble] unescape title
* [RUTV] Fix format sorting by [Lesmiscore](https://github.com/Lesmiscore)
* [ruutu] Detect embeds by [tpikonen](https://github.com/tpikonen)
* [tenplay] Improve extractor by [aarubui](https://github.com/aarubui)
* [TikTok] Fix URLs with user id by [hatienl0i261299](https://github.com/hatienl0i261299)
* [TikTokVM] Fix redirect to user URL
* [TVer] Fix extractor by [Lesmiscore](https://github.com/Lesmiscore)
* [TVer] Support landing page by [vvto33](https://github.com/vvto33)
* [twitcasting] Don't return multi_video for archive with single hls manifest by [Lesmiscore](https://github.com/Lesmiscore)
* [veo] Fix `_VALID_URL`
* [Veo] Fix extractor by [i6t](https://github.com/i6t)
* [viki] Don't attempt to modify URLs with signature by [nyuszika7h](https://github.com/nyuszika7h)
* [viu] Fix bypass for preview by [zackmark29](https://github.com/zackmark29)
* [viu] Fixed extractor by [zackmark29](https://github.com/zackmark29), [pukkandan](https://github.com/pukkandan)
* [web.archive:youtube] Make CDX API requests non-fatal by [coletdjnz](https://github.com/coletdjnz)
* [wget] Fix proxy by [kikuyan](https://github.com/kikuyan), [coletdjnz](https://github.com/coletdjnz)
* [xnxx] Add `xnxx3.com` by [rozari0](https://github.com/rozari0)
* [youtube] **Add new age-gate bypass** by [zerodytrash](https://github.com/zerodytrash), [pukkandan](https://github.com/pukkandan)
* [youtube] Add extractor-arg to skip auto-translated subs
* [youtube] Avoid false positives when detecting damaged formats
* [youtube] Detect DRM better by [shirt](https://github.com/shirt-dev)
* [youtube] Fix auto-translated automatic captions
* [youtube] Fix pagination of `membership` tab
* [youtube] Fix uploader for collaborative playlists by [coletdjnz](https://github.com/coletdjnz)
* [youtube] Improve video upload date handling by [coletdjnz](https://github.com/coletdjnz)
* [youtube:api] Prefer minified JSON response by [coletdjnz](https://github.com/coletdjnz)
* [youtube:search] Support hashtag entries by [coletdjnz](https://github.com/coletdjnz)
* [youtube:tab] Fix duration extraction for shorts by [coletdjnz](https://github.com/coletdjnz)
* [youtube:tab] Minor improvements
* [youtube:tab] Return shorts url if video is a short by [coletdjnz](https://github.com/coletdjnz)
* [Zattoo] Fix extractors by [goggle](https://github.com/goggle)
* [Zingmp3] Fix signature by [hatienl0i261299](https://github.com/hatienl0i261299)


### 2022.03.08.1

* [cleanup] Refactor `__init__.py`
* [build] Fix bug

### 2022.03.08

* Merge youtube-dl: Upto [commit/6508688](https://github.com/ytdl-org/youtube-dl/commit/6508688e88c83bb811653083db9351702cd39a6a) (except NDR)
* Add regex operator and quoting to format filters by [lukasfink1](https://github.com/lukasfink1)
* Add brotli content-encoding support by [coletdjnz](https://github.com/coletdjnz)
* Add pre-processor stage `after_filter`
* Better error message when no `--live-from-start` format
* Create necessary directories for `--print-to-file`
* Fill more fields for playlists by [Lesmiscore](https://github.com/Lesmiscore)
* Fix `-all` for `--sub-langs`
* Fix doubling of `video_id` in `ExtractorError`
* Fix for when stdout/stderr encoding is `None`
* Handle negative duration from extractor
* Implement `--add-header` without modifying `std_headers`
* Obey `--abort-on-error` for "ffmpeg not installed"
* Set `webpage_url_...` from `webpage_url` and not input URL
* Tolerate failure to `--write-link` due to unknown URL
* [aria2c] Add `--http-accept-gzip=true`
* [build] Update pyinstaller to 4.10 by [shirt](https://github.com/shirt-dev)
* [cookies] Update MacOS12 `Cookies.binarycookies` location by [mdpauley](https://github.com/mdpauley)
* [devscripts] Improve `prepare_manpage`
* [downloader] Do not use aria2c for non-native `m3u8`
* [downloader] Obey `--file-access-retries` when deleting/renaming by [ehoogeveen-medweb](https://github.com/ehoogeveen-medweb)
* [extractor] Allow `http_headers` to be specified for `thumbnails`
* [extractor] Extract subtitles from manifests for vimeo, globo, kaltura, svt by [fstirlitz](https://github.com/fstirlitz)
* [extractor] Fix for manifests without period duration by [dirkf](https://github.com/dirkf), [pukkandan](https://github.com/pukkandan)
* [extractor] Support `--mark-watched` without `_NETRC_MACHINE` by [coletdjnz](https://github.com/coletdjnz)
* [FFmpegConcat] Abort on `--simulate`
* [FormatSort] Consider `acodec`=`ogg` as `vorbis`
* [fragment] Fix bugs around resuming with Range by [Lesmiscore](https://github.com/Lesmiscore)
* [fragment] Improve `--live-from-start` for YouTube livestreams by [Lesmiscore](https://github.com/Lesmiscore)
* [generic] Pass referer to extracted formats
* [generic] Set rss `guid` as video id by [Bricio](https://github.com/Bricio)
* [options] Better ambiguous option resolution
* [options] Rename `--clean-infojson` to `--clean-info-json`
* [SponsorBlock] Fixes for highlight and "full video labels" by [nihil-admirari](https://github.com/nihil-admirari)
* [Sponsorblock] minor fixes by [nihil-admirari](https://github.com/nihil-admirari)
* [utils] Better traceback for `ExtractorError`
* [utils] Fix file locking for AOSP by [jakeogh](https://github.com/jakeogh)
* [utils] Improve file locking
* [utils] OnDemandPagedList: Do not download pages after error
* [utils] render_table: Fix character calculation for removing extra gap by [Lesmiscore](https://github.com/Lesmiscore)
* [utils] Use `locked_file` for `sanitize_open` by [jakeogh](https://github.com/jakeogh)
* [utils] Validate `DateRange` input
* [utils] WebSockets wrapper for non-async functions by [Lesmiscore](https://github.com/Lesmiscore)
* [cleanup] Don't pass protocol to `_extract_m3u8_formats` for live videos
* [cleanup] Remove extractors for some dead websites by [marieell](https://github.com/marieell)
* [cleanup, docs] Misc cleanup
* [AbemaTV] Add extractors by [Lesmiscore](https://github.com/Lesmiscore)
* [adobepass] Add Suddenlink MSO by [CplPwnies](https://github.com/CplPwnies)
* [ant1newsgr] Add extractor by [zmousm](https://github.com/zmousm)
* [bigo] Add extractor by [Lesmiscore](https://github.com/Lesmiscore)
* [Caltrans] Add extractor by [Bricio](https://github.com/Bricio)
* [daystar] Add extractor by [hatienl0i261299](https://github.com/hatienl0i261299)
* [fc2:live] Add extractor by [Lesmiscore](https://github.com/Lesmiscore)
* [fptplay] Add extractor by [hatienl0i261299](https://github.com/hatienl0i261299)
* [murrtube] Add extractor by [cyberfox1691](https://github.com/cyberfox1691)
* [nfb] Add extractor by [ofkz](https://github.com/ofkz)
* [niconico] Add playlist extractors and refactor by [Lesmiscore](https://github.com/Lesmiscore)
* [peekvids] Add extractor by [schn0sch](https://github.com/schn0sch)
* [piapro] Add extractor by [pycabbage](https://github.com/pycabbage), [Lesmiscore](https://github.com/Lesmiscore)
* [rokfin] Add extractor by [P-reducible](https://github.com/P-reducible), [pukkandan](https://github.com/pukkandan)
* [rokfin] Add stack and channel extractors by [P-reducible](https://github.com/P-reducible), [pukkandan](https://github.com/pukkandan)
* [ruv.is] Add extractor by [iw0nderhow](https://github.com/iw0nderhow)
* [telegram] Add extractor by [hatienl0i261299](https://github.com/hatienl0i261299)
* [VideocampusSachsen] Add extractors by [FestplattenSchnitzel](https://github.com/FestplattenSchnitzel)
* [xinpianchang] Add extractor by [hatienl0i261299](https://github.com/hatienl0i261299)
* [abc] Support 1080p by [Ronnnny](https://github.com/Ronnnny)
* [afreecatv] Support password-protected livestreams by [wlritchi](https://github.com/wlritchi)
* [ard] Fix valid URL
* [ATVAt] Detect geo-restriction by [marieell](https://github.com/marieell)
* [bandcamp] Detect acodec
* [bandcamp] Fix user URLs by [lyz-code](https://github.com/lyz-code)
* [bbc] Fix extraction of news articles by [ajj8](https://github.com/ajj8)
* [beeg] Fix extractor by [Bricio](https://github.com/Bricio)
* [bigo] Fix extractor to not to use `form_params`
* [Bilibili] Pass referer for all formats by [blackgear](https://github.com/blackgear)
* [Biqle] Fix extractor by [Bricio](https://github.com/Bricio)
* [ccma] Fix timestamp parsing by [nyuszika7h](https://github.com/nyuszika7h)
* [crunchyroll] Better error reporting on login failure by [tejing1](https://github.com/tejing1)
* [cspan] Support of C-Span congress videos by [Grabien](https://github.com/Grabien)
* [dropbox] fix regex by [zenerdi0de](https://github.com/zenerdi0de)
* [fc2] Fix extraction by [Lesmiscore](https://github.com/Lesmiscore)
* [fujitv] Extract resolution for free sources by [YuenSzeHong](https://github.com/YuenSzeHong)
* [Gettr] Add `GettrStreamingIE` by [i6t](https://github.com/i6t)
* [Gettr] Fix formats order by [i6t](https://github.com/i6t)
* [Gettr] Improve extractor by [i6t](https://github.com/i6t)
* [globo] Expand valid URL by [Bricio](https://github.com/Bricio)
* [lbry] Fix `--ignore-no-formats-error`
* [manyvids] Extract `uploader` by [regarten](https://github.com/regarten)
* [mildom] Fix linter
* [mildom] Rework extractors by [Lesmiscore](https://github.com/Lesmiscore)
* [mirrativ] Cleanup extractor code by [Lesmiscore](https://github.com/Lesmiscore)
* [nhk] Add support for NHK for School by [Lesmiscore](https://github.com/Lesmiscore)
* [niconico:tag] Add support for searching tags
* [nrk] Add fallback API
* [peekvids] Use JSON-LD by [schn0sch](https://github.com/schn0sch)
* [peertube] Add media.fsfe.org by [mxmehl](https://github.com/mxmehl)
* [rtvs] Fix extractor by [Bricio](https://github.com/Bricio)
* [spiegel] Fix `_VALID_URL`
* [ThumbnailsConvertor] Support `webp`
* [tiktok] Fix `vm.tiktok`/`vt.tiktok` URLs
* [tubitv] Fix/improve TV series extraction by [bbepis](https://github.com/bbepis)
* [tumblr] Fix extractor by [foghawk](https://github.com/foghawk)
* [twitcasting] Add fallback for finding running live by [Lesmiscore](https://github.com/Lesmiscore)
* [TwitCasting] Check for password protection by [Lesmiscore](https://github.com/Lesmiscore)
* [twitcasting] Fix extraction by [Lesmiscore](https://github.com/Lesmiscore)
* [twitch] Fix field name of `view_count`
* [twitter] Fix for private videos by [iphoting](https://github.com/iphoting)
* [washingtonpost] Fix extractor by [Bricio](https://github.com/Bricio)
* [youtube:tab] Add `approximate_date` extractor-arg
* [youtube:tab] Follow redirect to regional channel  by [coletdjnz](https://github.com/coletdjnz)
* [youtube:tab] Reject webpage data if redirected to home page
* [youtube] De-prioritize potentially damaged formats
* [youtube] Differentiate descriptive audio by language code
* [youtube] Ensure subtitle urls are absolute by [coletdjnz](https://github.com/coletdjnz)
* [youtube] Escape possible `$` in `_extract_n_function_name` regex by [Lesmiscore](https://github.com/Lesmiscore)
* [youtube] Fix automatic captions
* [youtube] Fix n-sig extraction for phone player JS by [MinePlayersPE](https://github.com/MinePlayersPE)
* [youtube] Further de-prioritize 3gp format
* [youtube] Label original auto-subs
* [youtube] Prefer UTC upload date for videos by [coletdjnz](https://github.com/coletdjnz)
* [zaq1] Remove dead extractor by [marieell](https://github.com/marieell)
* [zee5] Support web-series by [Aniruddh-J](https://github.com/Aniruddh-J)
* [zingmp3] Fix extractor by [hatienl0i261299](https://github.com/hatienl0i261299)
* [zoom] Add support for screen cast by [Mipsters](https://github.com/Mipsters)


### 2022.02.04

* [youtube:search] Fix extractor by [coletdjnz](https://github.com/coletdjnz)
* [youtube:search] Add tests
* [twitcasting] Enforce UTF-8 for POST payload by [Lesmiscore](https://github.com/Lesmiscore)
* [mediaset] Fix extractor by [nixxo](https://github.com/nixxo)
* [websocket] Make syntax error in `websockets` module non-fatal

### 2022.02.03

* Merge youtube-dl: Upto [commit/78ce962](https://github.com/ytdl-org/youtube-dl/commit/78ce962f4fe020994c216dd2671546fbe58a5c67)
* Add option `--print-to-file`
* Make nested --config-locations relative to parent file
* Ensure `_type` is present in `info.json`
* Fix `--compat-options list-formats`
* Fix/improve `InAdvancePagedList`
* [downloader/ffmpeg] Handle unknown formats better
* [outtmpl] Handle `-o ""` better
* [outtmpl] Handle hard-coded file extension better
* [extractor] Add convenience function `_yes_playlist`
* [extractor] Allow non-fatal `title` extraction
* [extractor] Extract video inside `Article` json_ld
* [generic] Allow further processing of json_ld URL
* [cookies] Fix keyring selection for unsupported desktops
* [utils] Strip double spaces in `clean_html` by [dirkf](https://github.com/dirkf)
* [aes] Add `unpad_pkcs7`
* [test] Fix `test_youtube_playlist_noplaylist`
* [docs,cleanup] Misc cleanup
* [dplay] Add extractors for site changes by [Sipherdrakon](https://github.com/Sipherdrakon)
* [ertgr] Add  extractors by [zmousm](https://github.com/zmousm), [dirkf](https://github.com/dirkf)
* [Musicdex] Add extractors by [Ashish0804](https://github.com/Ashish0804)
* [YandexVideoPreview] Add extractor by [KiberInfinity](https://github.com/KiberInfinity)
* [youtube] Add extractor `YoutubeMusicSearchURLIE`
* [archive.org] Ignore unnecessary files
* [Bilibili] Add 8k support by [u-spec-png](https://github.com/u-spec-png)
* [bilibili] Fix extractor, make anthology title non-fatal
* [CAM4] Add thumbnail extraction by [alerikaisattera](https://github.com/alerikaisattera)
* [cctv] De-prioritize sample format
* [crunchyroll:beta] Add cookies support by [tejing1](https://github.com/tejing1)
* [crunchyroll] Fix login by [tejing1](https://github.com/tejing1)
* [doodstream] Fix extractor
* [fc2] Fix extraction by [Lesmiscore](https://github.com/Lesmiscore)
* [FFmpegConcat] Abort on --skip-download and download errors
* [Fujitv] Extract metadata and support premium by [YuenSzeHong](https://github.com/YuenSzeHong)
* [globo] Fix extractor by [Bricio](https://github.com/Bricio)
* [glomex] Simplify embed detection
* [GoogleSearch] Fix extractor
* [Instagram] Fix extraction when logged in by [MinePlayersPE](https://github.com/MinePlayersPE)
* [iq.com] Add VIP support by [MinePlayersPE](https://github.com/MinePlayersPE)
* [mildom] Fix extractor by [lazypete365](https://github.com/lazypete365)
* [MySpass] Fix video url processing by [trassshhub](https://github.com/trassshhub)
* [Odnoklassniki] Improve embedded players extraction by [KiberInfinity](https://github.com/KiberInfinity)
* [orf:tvthek] Lazy playlist extraction and obey --no-playlist
* [Pladform] Fix redirection to external player by [KiberInfinity](https://github.com/KiberInfinity)
* [ThisOldHouse] Improve Premium URL check by [Ashish0804](https://github.com/Ashish0804)
* [TikTok] Iterate through app versions by [MinePlayersPE](https://github.com/MinePlayersPE)
* [tumblr] Fix 403 errors and handle vimeo embeds by [foghawk](https://github.com/foghawk)
* [viki] Fix "Bad request" for manifest by [nyuszika7h](https://github.com/nyuszika7h)
* [Vimm] add recording extractor by [alerikaisattera](https://github.com/alerikaisattera)
* [web.archive:youtube] Add `ytarchive:` prefix and misc cleanup
* [youtube:api] Do not use seek when reading HTTPError response by [coletdjnz](https://github.com/coletdjnz)
* [youtube] Fix n-sig for player e06dea74
* [youtube, cleanup] Misc fixes and cleanup


### 2022.01.21

* Add option `--concat-playlist` to **concat videos in a playlist**
* Allow **multiple and nested configuration files**
* Add more post-processing stages (`after_video`, `playlist`)
* Allow `--exec` to be run at any post-processing stage (Deprecates `--exec-before-download`)
* Allow `--print` to be run at any post-processing stage
* Allow listing formats, thumbnails, subtitles using `--print` by [pukkandan](https://github.com/pukkandan), [Zirro](https://github.com/Zirro)
* Add fields `video_autonumber`, `modified_date`, `modified_timestamp`, `playlist_count`, `channel_follower_count`
* Add key `requested_downloads` in the root `info_dict`
* Write `download_archive` only after all formats are downloaded
* [FfmpegMetadata] Allow setting metadata of individual streams using `meta<n>_` prefix
* Add option `--legacy-server-connect` by [xtkoba](https://github.com/xtkoba)
* Allow escaped `,` in `--extractor-args`
* Allow unicode characters in `info.json`
* Check for existing thumbnail/subtitle in final directory
* Don't treat empty containers as `None` in `sanitize_info`
* Fix `-s --ignore-no-formats --force-write-archive`
* Fix live title for multiple formats
* List playlist thumbnails in `--list-thumbnails`
* Raise error if subtitle download fails
* [cookies] Fix bug when keyring is unspecified
* [ffmpeg] Ignore unknown streams, standardize use of `-map 0`
* [outtmpl] Alternate form for `D` and fix suffix's case
* [utils] Add `Sec-Fetch-Mode` to `std_headers`
* [utils] Fix `format_bytes` output for Bytes by [pukkandan](https://github.com/pukkandan), [mdawar](https://github.com/mdawar)
* [utils] Handle `ss:xxx` in `parse_duration`
* [utils] Improve parsing for nested HTML elements by [zmousm](https://github.com/zmousm), [pukkandan](https://github.com/pukkandan)
* [utils] Use key `None` in `traverse_obj` to return as-is
* [extractor] Detect more subtitle codecs in MPD manifests by [fstirlitz](https://github.com/fstirlitz)
* [extractor] Extract chapters from JSON-LD by [iw0nderhow](https://github.com/iw0nderhow), [pukkandan](https://github.com/pukkandan)
* [extractor] Extract thumbnails from JSON-LD by [nixxo](https://github.com/nixxo)
* [extractor] Improve `url_result` and related
* [generic] Improve KVS player extraction by [trassshhub](https://github.com/trassshhub)
* [build] Reduce dependency on third party workflows
* [extractor,cleanup] Use `_search_nextjs_data`, `format_field`
* [cleanup] Minor fixes and cleanup
* [docs] Improvements
* [test] Fix TestVerboseOutput
* [afreecatv] Add livestreams extractor by [wlritchi](https://github.com/wlritchi)
* [callin] Add extractor by [foghawk](https://github.com/foghawk)
* [CrowdBunker] Add extractors by [Ashish0804](https://github.com/Ashish0804)
* [daftsex] Add extractors by [k3ns1n](https://github.com/k3ns1n)
* [digitalconcerthall] Add extractor by [teridon](https://github.com/teridon)
* [Drooble] Add extractor by [u-spec-png](https://github.com/u-spec-png)
* [EuropeanTour] Add extractor by [Ashish0804](https://github.com/Ashish0804)
* [iq.com] Add extractors by [MinePlayersPE](https://github.com/MinePlayersPE)
* [KelbyOne] Add extractor by [Ashish0804](https://github.com/Ashish0804)
* [LnkIE] Add extractor by [Ashish0804](https://github.com/Ashish0804)
* [MainStreaming] Add extractor by [coletdjnz](https://github.com/coletdjnz)
* [megatvcom] Add extractors by [zmousm](https://github.com/zmousm)
* [Newsy] Add extractor by [Ashish0804](https://github.com/Ashish0804)
* [noodlemagazine] Add extractor by [trassshhub](https://github.com/trassshhub)
* [PokerGo] Add extractors by [Ashish0804](https://github.com/Ashish0804)
* [Pornez] Add extractor by [mozlima](https://github.com/mozlima)
* [PRX] Add Extractors by [coletdjnz](https://github.com/coletdjnz)
* [RTNews] Add extractor by [Ashish0804](https://github.com/Ashish0804)
* [Rule34video] Add extractor by [trassshhub](https://github.com/trassshhub)
* [tvopengr] Add extractors by [zmousm](https://github.com/zmousm)
* [Vimm] Add extractor by [alerikaisattera](https://github.com/alerikaisattera)
* [glomex] Add extractors by [zmousm](https://github.com/zmousm)
* [instagram] Add story/highlight extractor by [u-spec-png](https://github.com/u-spec-png)
* [openrec] Add movie extractor by [Lesmiscore](https://github.com/Lesmiscore)
* [rai] Add Raiplaysound extractors by [nixxo](https://github.com/nixxo), [pukkandan](https://github.com/pukkandan)
* [aparat] Fix extractor
* [ard] Extract subtitles by [fstirlitz](https://github.com/fstirlitz)
* [BiliIntl] Add login by [MinePlayersPE](https://github.com/MinePlayersPE)
* [CeskaTelevize] Use `http` for manifests
* [CTVNewsIE] Add fallback for video search by [Ashish0804](https://github.com/Ashish0804)
* [dplay] Migrate DiscoveryPlusItaly to DiscoveryPlus by [timendum](https://github.com/timendum)
* [dplay] Re-structure DiscoveryPlus extractors
* [Dropbox] Support password protected files and more formats by [zenerdi0de](https://github.com/zenerdi0de)
* [facebook] Fix extraction from groups
* [facebook] Improve title and uploader extraction
* [facebook] Parse dash manifests
* [fox] Extract m3u8 from preview by [ischmidt20](https://github.com/ischmidt20)
* [funk] Support origin URLs
* [gfycat] Fix `uploader`
* [gfycat] Support embeds by [coletdjnz](https://github.com/coletdjnz)
* [hotstar] Add extractor args to ignore tags by [Ashish0804](https://github.com/Ashish0804)
* [hrfernsehen] Fix ardloader extraction by [CreaValix](https://github.com/CreaValix)
* [instagram] Fix username extraction for stories and highlights by [nyuszika7h](https://github.com/nyuszika7h)
* [kakao] Detect geo-restriction
* [line] Remove `tv.line.me` by [sian1468](https://github.com/sian1468)
* [mixch] Add `MixchArchiveIE` by [Lesmiscore](https://github.com/Lesmiscore)
* [mixcloud] Detect restrictions by [llacb47](https://github.com/llacb47)
* [NBCSports] Fix extraction of platform URLs by [ischmidt20](https://github.com/ischmidt20)
* [Nexx] Extract more metadata by [MinePlayersPE](https://github.com/MinePlayersPE)
* [Nexx] Support 3q CDN by [MinePlayersPE](https://github.com/MinePlayersPE)
* [pbs] de-prioritize AD formats
* [PornHub,YouTube] Refresh onion addresses by [unit193](https://github.com/unit193)
* [RedBullTV] Parse subtitles from manifest by [Ashish0804](https://github.com/Ashish0804)
* [streamcz] Fix extractor by [arkamar](https://github.com/arkamar), [pukkandan](https://github.com/pukkandan)
* [Ted] Rewrite extractor by [pukkandan](https://github.com/pukkandan), [trassshhub](https://github.com/trassshhub)
* [Theta] Fix valid URL by [alerikaisattera](https://github.com/alerikaisattera)
* [ThisOldHouseIE] Add support for premium videos by [Ashish0804](https://github.com/Ashish0804)
* [TikTok] Fix extraction for sigi-based webpages, add API fallback by [MinePlayersPE](https://github.com/MinePlayersPE)
* [TikTok] Pass cookies to formats, and misc fixes by [MinePlayersPE](https://github.com/MinePlayersPE)
* [TikTok] Extract captions, user thumbnail by [MinePlayersPE](https://github.com/MinePlayersPE)
* [TikTok] Change app version by [MinePlayersPE](https://github.com/MinePlayersPE), [llacb47](https://github.com/llacb47)
* [TVer] Extract message for unaired live by [Lesmiscore](https://github.com/Lesmiscore)
* [twitcasting] Refactor extractor by [Lesmiscore](https://github.com/Lesmiscore)
* [twitter] Fix video in quoted tweets
* [veoh] Improve extractor by [foghawk](https://github.com/foghawk)
* [vk] Capture `clip` URLs
* [vk] Fix VKUserVideosIE by [Ashish0804](https://github.com/Ashish0804)
* [vk] Improve `_VALID_URL` by [k3ns1n](https://github.com/k3ns1n)
* [VrtNU] Handle empty title by [pgaig](https://github.com/pgaig)
* [XVideos] Check HLS formats by [MinePlayersPE](https://github.com/MinePlayersPE)
* [yahoo:gyao] Improved playlist handling by [hyano](https://github.com/hyano)
* [youtube:tab] Extract more playlist metadata by [coletdjnz](https://github.com/coletdjnz), [pukkandan](https://github.com/pukkandan)
* [youtube:tab] Raise error on tab redirect by [krichbanana](https://github.com/krichbanana), [coletdjnz](https://github.com/coletdjnz)
* [youtube] Update Innertube clients by [coletdjnz](https://github.com/coletdjnz)
* [youtube] Detect live-stream embeds
* [youtube] Do not return `upload_date` for playlists
* [youtube] Extract channel subscriber count by [coletdjnz](https://github.com/coletdjnz)
* [youtube] Make invalid storyboard URL non-fatal
* [youtube] Enforce UTC, update innertube clients and tests by [coletdjnz](https://github.com/coletdjnz)
* [zdf] Add chapter extraction by [iw0nderhow](https://github.com/iw0nderhow)
* [zee5] Add geo-bypass


### 2021.12.27

* Avoid recursion error when re-extracting info
* [ffmpeg] Fix position of `--ppa`
* [aria2c] Don't show progress when `--no-progress`
* [cookies] Support other keyrings by [mbway](https://github.com/mbway)
* [EmbedThumbnail] Prefer AtomicParsley over ffmpeg if available
* [generic] Fix HTTP KVS Player by [git-anony-mouse](https://github.com/git-anony-mouse)
* [ThumbnailsConvertor] Fix for when there are no thumbnails
* [docs] Add examples for using `TYPES:` in `-P`/`-o`
* [PixivSketch] Add extractors by [nao20010128nao](https://github.com/nao20010128nao)
* [tiktok] Add music, sticker and tag IEs by [MinePlayersPE](https://github.com/MinePlayersPE)
* [BiliIntl] Fix extractor by [MinePlayersPE](https://github.com/MinePlayersPE)
* [CBC] Fix URL regex
* [tiktok] Fix `extractor_key` used in archive
* [youtube] **End `live-from-start` properly when stream ends with 403**
* [Zee5] Fix VALID_URL for tv-shows by [Ashish0804](https://github.com/Ashish0804)

### 2021.12.25

* [dash,youtube] **Download live from start to end** by [nao20010128nao](https://github.com/nao20010128nao), [pukkandan](https://github.com/pukkandan)
    * Add option `--live-from-start` to enable downloading live videos from start
    * Add key `is_from_start` in formats to identify formats (of live videos) that downloads from start
    * [dash] Create protocol `http_dash_segments_generator` that allows a function to be passed instead of fragments
    * [fragment] Allow multiple live dash formats to download simultaneously
    * [youtube] Implement fragment re-fetching for the live dash formats
    * [youtube] Re-extract dash manifest every 5 hours (manifest expires in 6hrs)
    * [postprocessor/ffmpeg] Add `FFmpegFixupDuplicateMoovPP` to fixup duplicated moov atoms
    * Known issues:
        * Ctrl+C doesn't work on Windows when downloading multiple formats
        * If video becomes private, download hangs
* [SponsorBlock] Add `Filler` and `Highlight` categories by [nihil-admirari](https://github.com/nihil-admirari), [pukkandan](https://github.com/pukkandan)
    * Change `--sponsorblock-cut all` to `--sponsorblock-cut default` if you do not want filler sections to be removed
* Add field `webpage_url_domain`
* Add interactive format selection with `-f -`
* Add option `--file-access-retries` by [ehoogeveen-medweb](https://github.com/ehoogeveen-medweb)
* [outtmpl] Add alternate forms `S`, `D` and improve `id` detection
* [outtmpl] Add operator `&` for replacement text by [PilzAdam](https://github.com/PilzAdam)
* [EmbedSubtitle] Disable duration check temporarily
* [extractor] Add `_search_nuxt_data` by [nao20010128nao](https://github.com/nao20010128nao)
* [extractor] Ignore errors in comment extraction when `-i` is given
* [extractor] Standardize `_live_title`
* [FormatSort] Prevent incorrect deprecation warning
* [generic] Extract m3u8 formats from JSON-LD
* [postprocessor/ffmpeg] Always add `faststart`
* [utils] Fix parsing `YYYYMMDD` dates in Nov/Dec by [wlritchi](https://github.com/wlritchi)
* [utils] Improve `parse_count`
* [utils] Update `std_headers` by [kikuyan](https://github.com/kikuyan), [fstirlitz](https://github.com/fstirlitz)
* [lazy_extractors] Fix for search IEs
* [extractor] Support default implicit graph in JSON-LD by [zmousm](https://github.com/zmousm)
* Allow `--no-write-thumbnail` to override `--write-all-thumbnail`
* Fix `--throttled-rate`
* Fix control characters being printed to `--console-title`
* Fix PostProcessor hooks not registered for some PPs
* Pre-process when using `--flat-playlist`
* Remove known invalid thumbnails from `info_dict`
* Add warning when using `-f best`
* Use `parse_duration` for `--wait-for-video` and some minor fix
* [test/download] Add more fields
* [test/download] Ignore field `webpage_url_domain` by [std-move](https://github.com/std-move)
* [compat] Suppress errors in enabling VT mode
* [docs] Improve manpage format by [iw0nderhow](https://github.com/iw0nderhow), [pukkandan](https://github.com/pukkandan)
* [docs,cleanup] Minor fixes and cleanup
* [cleanup] Fix some typos by [unit193](https://github.com/unit193)
* [ABC:iview] Add show extractor by [pabs3](https://github.com/pabs3)
* [dropout] Add extractor by [TwoThousandHedgehogs](https://github.com/TwoThousandHedgehogs), [pukkandan](https://github.com/pukkandan)
* [GameJolt] Add extractors by [MinePlayersPE](https://github.com/MinePlayersPE)
* [gofile] Add extractor by [Jertzukka](https://github.com/Jertzukka), [Ashish0804](https://github.com/Ashish0804)
* [hse] Add extractors by [cypheron](https://github.com/cypheron), [pukkandan](https://github.com/pukkandan)
* [NateTV] Add NateIE and NateProgramIE by [Ashish0804](https://github.com/Ashish0804), [Hyeeji](https://github.com/Hyeeji)
* [OpenCast] Add extractors by [bwildenhain](https://github.com/bwildenhain), [C0D3D3V](https://github.com/C0D3D3V)
* [rtve] Add `RTVEAudioIE` by [kebianizao](https://github.com/kebianizao)
* [Rutube] Add RutubeChannelIE by [Ashish0804](https://github.com/Ashish0804)
* [skeb] Add extractor by [nao20010128nao](https://github.com/nao20010128nao)
* [soundcloud] Add related tracks extractor by [Lapin0t](https://github.com/Lapin0t)
* [toggo] Add extractor by [nyuszika7h](https://github.com/nyuszika7h)
* [TrueID] Add extractor by [MinePlayersPE](https://github.com/MinePlayersPE)
* [audiomack] Update album and song VALID_URL by [abdullah-if](https://github.com/abdullah-if), [dirkf](https://github.com/dirkf)
* [CBC Gem] Extract 1080p formats by [DavidSkrundz](https://github.com/DavidSkrundz)
* [ceskatelevize] Fetch iframe from nextJS data by [mkubecek](https://github.com/mkubecek)
* [crackle] Look for non-DRM formats by [raleeper](https://github.com/raleeper)
* [dplay] Temporary fix for `discoveryplus.com/it`
* [DiscoveryPlusShowBaseIE] yield actual video id by [Ashish0804](https://github.com/Ashish0804)
* [Facebook] Handle redirect URLs
* [fujitv] Extract 1080p from `tv_android` m3u8 by [YuenSzeHong](https://github.com/YuenSzeHong)
* [gronkh] Support new URL pattern by [Sematre](https://github.com/Sematre)
* [instagram] Expand valid URL by [u-spec-png](https://github.com/u-spec-png)
* [Instagram] Try bypassing login wall with embed page by [MinePlayersPE](https://github.com/MinePlayersPE)
* [Jamendo] Fix use of `_VALID_URL_RE` by [jaller94](https://github.com/jaller94)
* [LBRY] Support livestreams by [Ashish0804](https://github.com/Ashish0804), [pukkandan](https://github.com/pukkandan)
* [NJPWWorld] Extract formats from m3u8 by [aarubui](https://github.com/aarubui)
* [NovaEmbed] update player regex by [std-move](https://github.com/std-move)
* [npr] Make SMIL extraction non-fatal by [r5d](https://github.com/r5d)
* [ntvcojp] Extract NUXT data by [nao20010128nao](https://github.com/nao20010128nao)
* [ok.ru] add mobile fallback by [nao20010128nao](https://github.com/nao20010128nao)
* [olympics] Add uploader and cleanup by [u-spec-png](https://github.com/u-spec-png)
* [ondemandkorea] Update `jw_config` regex by [julien-hadleyjack](https://github.com/julien-hadleyjack)
* [PlutoTV] Expand `_VALID_URL`
* [RaiNews] Fix extractor by [nixxo](https://github.com/nixxo)
* [RCTIPlusSeries] Lazy extraction and video type selection by [MinePlayersPE](https://github.com/MinePlayersPE)
* [redtube] Handle formats delivered inside a JSON by [dirkf](https://github.com/dirkf), [nixxo](https://github.com/nixxo)
* [SonyLiv] Add OTP login support by [Ashish0804](https://github.com/Ashish0804)
* [Steam] Fix extractor by [u-spec-png](https://github.com/u-spec-png)
* [TikTok] Pass cookies to mobile API by [MinePlayersPE](https://github.com/MinePlayersPE)
* [trovo] Fix inheritance of `TrovoChannelBaseIE`
* [TVer] Extract better thumbnails by [YuenSzeHong](https://github.com/YuenSzeHong)
* [vimeo] Extract chapters
* [web.archive:youtube] Improve metadata extraction by [coletdjnz](https://github.com/coletdjnz)
* [youtube:comments] Add more options for limiting number of comments extracted by [coletdjnz](https://github.com/coletdjnz)
* [youtube:tab] Extract more metadata from feeds/channels/playlists by [coletdjnz](https://github.com/coletdjnz)
* [youtube:tab] Extract video thumbnails from playlist by [coletdjnz](https://github.com/coletdjnz), [pukkandan](https://github.com/pukkandan)
* [youtube:tab] Ignore query when redirecting channel to playlist and cleanup of related code
* [youtube] Fix `ytsearchdate`
* [zdf] Support videos with different ptmd location by [iw0nderhow](https://github.com/iw0nderhow)
* [zee5] Support /episodes in URL


### 2021.12.01

* **Add option `--wait-for-video` to wait for scheduled streams**
* Add option `--break-per-input` to apply --break-on... to each input URL
* Add option `--embed-info-json` to embed info.json in mkv
* Add compat-option `embed-metadata`
* Allow using a custom format selector through API
* [AES] Add ECB mode by [nao20010128nao](https://github.com/nao20010128nao)
* [build] Fix MacOS Build
* [build] Save Git HEAD at release alongside version info
* [build] Use `workflow_dispatch` for release
* [downloader/ffmpeg] Fix for direct videos inside mpd manifests
* [downloader] Add colors to download progress
* [EmbedSubtitles] Slightly relax duration check and related cleanup
* [ExtractAudio] Fix conversion to `wav` and `vorbis`
* [ExtractAudio] Support `alac`
* [extractor] Extract `average_rating` from JSON-LD
* [FixupM3u8] Fixup MPEG-TS in MP4 container
* [generic] Support mpd manifests without extension by [shirt](https://github.com/shirt-dev)
* [hls] Better FairPlay DRM detection by [nyuszika7h](https://github.com/nyuszika7h)
* [jsinterp] Fix splice to handle float (for youtube js player f1ca6900)
* [utils] Allow alignment in `render_table` and add tests
* [utils] Fix `PagedList`
* [utils] Fix error when copying `LazyList`
* Clarify video/audio-only formats in -F
* Ensure directory exists when checking formats
* Ensure path for link files exists by [Zirro](https://github.com/Zirro)
* Ensure same config file is not loaded multiple times
* Fix `postprocessor_hooks`
* Fix `--break-on-archive` when pre-checking
* Fix `--check-formats` for `mhtml`
* Fix `--load-info-json` of playlists with failed entries
* Fix `--trim-filename` when filename has `.`
* Fix bug in parsing `--add-header`
* Fix error in `report_unplayable_conflict` by [shirt](https://github.com/shirt-dev)
* Fix writing playlist infojson with `--no-clean-infojson`
* Validate --get-bypass-country
* [blogger] Add extractor by [pabs3](https://github.com/pabs3)
* [breitbart] Add extractor by [Grabien](https://github.com/Grabien)
* [CableAV] Add extractor by [j54vc1bk](https://github.com/j54vc1bk)
* [CanalAlpha] Add extractor by [Ashish0804](https://github.com/Ashish0804)
* [CozyTV] Add extractor by [Ashish0804](https://github.com/Ashish0804)
* [CPTwentyFour] Add extractor by [Ashish0804](https://github.com/Ashish0804)
* [DiscoveryPlus] Add `DiscoveryPlusItalyShowIE` by [Ashish0804](https://github.com/Ashish0804)
* [ESPNCricInfo] Add extractor by [Ashish0804](https://github.com/Ashish0804)
* [LinkedIn] Add extractor by [u-spec-png](https://github.com/u-spec-png)
* [mixch] Add extractor by [nao20010128nao](https://github.com/nao20010128nao)
* [nebula] Add `NebulaCollectionIE` and rewrite extractor by [hheimbuerger](https://github.com/hheimbuerger)
* [OneFootball] Add extractor by [Ashish0804](https://github.com/Ashish0804)
* [peer.tv] Add extractor by [u-spec-png](https://github.com/u-spec-png)
* [radiozet] Add extractor by [0xA7404A](https://github.com/0xA7404A) (Aurora)
* [redgifs] Add extractor by [chio0hai](https://github.com/chio0hai)
* [RedGifs] Add Search and User extractors by [Deer-Spangle](https://github.com/Deer-Spangle)
* [rtrfm] Add extractor by [pabs3](https://github.com/pabs3)
* [Streamff] Add extractor by [cntrl-s](https://github.com/cntrl-s)
* [Stripchat] Add extractor by [zulaport](https://github.com/zulaport)
* [Aljazeera] Fix extractor by [u-spec-png](https://github.com/u-spec-png)
* [AmazonStoreIE] Fix regex to not match vdp urls by [Ashish0804](https://github.com/Ashish0804)
* [ARDBetaMediathek] Handle new URLs
* [bbc] Get all available formats by [nyuszika7h](https://github.com/nyuszika7h)
* [Bilibili] Fix title extraction by [u-spec-png](https://github.com/u-spec-png)
* [CBC Gem] Fix for shows that don't have all seasons by [makeworld-the-better-one](https://github.com/makeworld-the-better-one)
* [curiositystream] Add more metadata
* [CuriosityStream] Fix series
* [DiscoveryPlus] Rewrite extractors by [Ashish0804](https://github.com/Ashish0804), [pukkandan](https://github.com/pukkandan)
* [HotStar] Set language field from tags by [Ashish0804](https://github.com/Ashish0804)
* [instagram, cleanup] Refactor extractors
* [Instagram] Display more login errors by [MinePlayersPE](https://github.com/MinePlayersPE)
* [itv] Fix extractor by [staubichsauger](https://github.com/staubichsauger), [pukkandan](https://github.com/pukkandan)
* [mediaklikk] Expand valid URL
* [MTV] Improve mgid extraction by [Sipherdrakon](https://github.com/Sipherdrakon), [kikuyan](https://github.com/kikuyan)
* [nexx] Better error message for unsupported format
* [NovaEmbed] Fix extractor by [pukkandan](https://github.com/pukkandan), [std-move](https://github.com/std-move)
* [PatreonUser] Do not capture RSS URLs
* [Reddit] Add support for 1080p videos by [xenova](https://github.com/xenova)
* [RoosterTeethSeries] Fix for multiple pages by [MinePlayersPE](https://github.com/MinePlayersPE)
* [sbs] Fix for movies and livestreams
* [Senate.gov] Add SenateGovIE and fix SenateISVPIE by [Grabien](https://github.com/Grabien), [pukkandan](https://github.com/pukkandan)
* [soundcloud:search] Fix pagination
* [tiktok:user] Set `webpage_url` correctly
* [Tokentube] Fix description by [u-spec-png](https://github.com/u-spec-png)
* [trovo] Fix extractor by [nyuszika7h](https://github.com/nyuszika7h)
* [tv2] Expand valid URL
* [Tvplayhome] Fix extractor by [pukkandan](https://github.com/pukkandan), [18928172992817182](https://github.com/18928172992817182)
* [Twitch:vod] Add chapters by [mpeter50](https://github.com/mpeter50)
* [twitch:vod] Extract live status by [DEvmIb](https://github.com/DEvmIb)
* [VidLii] Add 720p support by [mrpapersonic](https://github.com/mrpapersonic)
* [vimeo] Add fallback for config URL
* [vimeo] Sort http formats higher
* [WDR] Expand valid URL
* [willow] Add extractor by [aarubui](https://github.com/aarubui)
* [xvideos] Detect embed URLs by [4a1e2y5](https://github.com/4a1e2y5)
* [xvideos] Fix extractor by [Yakabuff](https://github.com/Yakabuff)
* [youtube, cleanup] Reorganize Tab and Search extractor inheritances
* [youtube:search_url] Add playlist/channel support
* [youtube] Add `default` player client by [coletdjnz](https://github.com/coletdjnz)
* [youtube] Add storyboard formats
* [youtube] Decrypt n-sig for URLs with `ratebypass`
* [youtube] Minor improvement to format sorting
* [cleanup] Add deprecation warnings
* [cleanup] Refactor `JSInterpreter._seperate`
* [Cleanup] Remove some unnecessary groups in regexes by [Ashish0804](https://github.com/Ashish0804)
* [cleanup] Misc cleanup


### 2021.11.10.1

* Temporarily disable MacOS Build

### 2021.11.10

* [youtube] **Fix throttling by decrypting n-sig**
* Merging extractors from [haruhi-dl](https://git.sakamoto.pl/laudom/haruhi-dl) by [selfisekai](https://github.com/selfisekai)
    * [extractor] Add `_search_nextjs_data`
    * [tvp] Fix extractors
    * [tvp] Add TVPStreamIE
    * [wppilot] Add extractors
    * [polskieradio] Add extractors
    * [radiokapital] Add extractors
    * [polsatgo] Add extractor by [selfisekai](https://github.com/selfisekai), [sdomi](https://github.com/sdomi)
* Separate `--check-all-formats` from `--check-formats`
* Approximate filesize from bitrate
* Don't create console in `windows_enable_vt_mode`
* Fix bug in `--load-infojson` of playlists
* [minicurses] Add colors to `-F` and standardize color-printing code
* [outtmpl] Add type `link` for internet shortcut files
* [outtmpl] Add alternate forms for `q` and `j`
* [outtmpl] Do not traverse `None`
* [fragment] Fix progress display in fragmented downloads
* [downloader/ffmpeg] Fix vtt download with ffmpeg
* [ffmpeg] Detect presence of setts and libavformat version
* [ExtractAudio] Rescale `--audio-quality` correctly by [CrypticSignal](https://github.com/CrypticSignal), [pukkandan](https://github.com/pukkandan)
* [ExtractAudio] Use `libfdk_aac` if available by [CrypticSignal](https://github.com/CrypticSignal)
* [FormatSort] `eac3` is better than `ac3`
* [FormatSort] Fix some fields' defaults
* [generic] Detect more json_ld
* [generic] parse jwplayer with only the json URL
* [extractor] Add keyword automatically to SearchIE descriptions
* [extractor] Fix some errors being converted to `ExtractorError`
* [utils] Add `join_nonempty`
* [utils] Add `jwt_decode_hs256` by [Ashish0804](https://github.com/Ashish0804)
* [utils] Create `DownloadCancelled` exception
* [utils] Parse `vp09` as vp9
* [utils] Sanitize URL when determining protocol
* [test/download] Fallback test to `bv`
* [docs] Minor documentation improvements
* [cleanup] Improvements to error and debug messages
* [cleanup] Minor fixes and cleanup
* [3speak] Add extractors by [Ashish0804](https://github.com/Ashish0804)
* [AmazonStore] Add extractor by [Ashish0804](https://github.com/Ashish0804)
* [Gab] Add extractor by [u-spec-png](https://github.com/u-spec-png)
* [mediaset] Add playlist support by [nixxo](https://github.com/nixxo)
* [MLSScoccer] Add extractor by [Ashish0804](https://github.com/Ashish0804)
* [N1] Add support for nova.rs by [u-spec-png](https://github.com/u-spec-png)
* [PlanetMarathi] Add extractor by [Ashish0804](https://github.com/Ashish0804)
* [RaiplayRadio] Add extractors by [frafra](https://github.com/frafra)
* [roosterteeth] Add series extractor
* [sky] Add `SkyNewsStoryIE` by [ajj8](https://github.com/ajj8)
* [youtube] Fix sorting for some videos
* [youtube] Populate `thumbnail` with the best "known" thumbnail
* [youtube] Refactor itag processing
* [youtube] Remove unnecessary no-playlist warning
* [youtube:tab] Add Invidious list for playlists/channels by [rhendric](https://github.com/rhendric)
* [Bilibili:comments] Fix infinite loop by [u-spec-png](https://github.com/u-spec-png)
* [ceskatelevize] Fix extractor by [flashdagger](https://github.com/flashdagger)
* [Coub] Fix media format identification by [wlritchi](https://github.com/wlritchi)
* [crunchyroll] Add extractor-args `language` and `hardsub`
* [DiscoveryPlus] Allow language codes in URL
* [imdb] Fix thumbnail by [ozburo](https://github.com/ozburo)
* [instagram] Add IOS URL support by [u-spec-png](https://github.com/u-spec-png)
* [instagram] Improve login code by [u-spec-png](https://github.com/u-spec-png)
* [Instagram] Improve metadata extraction by [u-spec-png](https://github.com/u-spec-png)
* [iPrima] Fix extractor by [stanoarn](https://github.com/stanoarn)
* [itv] Add support for ITV News by [ajj8](https://github.com/ajj8)
* [la7] Fix extractor by [nixxo](https://github.com/nixxo)
* [linkedin] Don't login multiple times
* [mtv] Fix some videos by [Sipherdrakon](https://github.com/Sipherdrakon)
* [Newgrounds] Fix description by [u-spec-png](https://github.com/u-spec-png)
* [Nrk] Minor fixes by [fractalf](https://github.com/fractalf)
* [Olympics] Fix extractor by [u-spec-png](https://github.com/u-spec-png)
* [piksel] Fix sorting
* [twitter] Do not sort by codec
* [viewlift] Add cookie-based login and series support by [Ashish0804](https://github.com/Ashish0804), [pukkandan](https://github.com/pukkandan)
* [vimeo] Detect source extension and misc cleanup by [flashdagger](https://github.com/flashdagger)
* [vimeo] Fix ondemand videos and direct URLs with hash
* [vk] Fix login and add subtitles by [kaz-us](https://github.com/kaz-us)
* [VLive] Add upload_date and thumbnail by [Ashish0804](https://github.com/Ashish0804)
* [VRT] Fix login by [pgaig](https://github.com/pgaig)
* [Vupload] Fix extractor by [u-spec-png](https://github.com/u-spec-png)
* [wakanim] Add support for MPD manifests by [nyuszika7h](https://github.com/nyuszika7h)
* [wakanim] Detect geo-restriction by [nyuszika7h](https://github.com/nyuszika7h)
* [ZenYandex] Fix extractor by [u-spec-png](https://github.com/u-spec-png)


### 2021.10.22

* [build] Improvements
    * Build standalone MacOS packages by [smplayer-dev](https://github.com/smplayer-dev)
    * Release windows exe built with `py2exe`
    * Enable lazy-extractors in releases
        * Set env var `YTDLP_NO_LAZY_EXTRACTORS` to forcefully disable this (experimental)
    * Clean up error reporting in update
    * Refactor `pyinst.py`, misc cleanup and improve docs
* [docs] Migrate issues to use forms by [Ashish0804](https://github.com/Ashish0804)
* [downloader] **Fix slow progress hooks**
    * This was causing HLS/DASH downloads to be extremely slow in some situations
* [downloader/ffmpeg] Improve simultaneous download and merge
* [EmbedMetadata] Allow overwriting all default metadata with `meta_default` key
* [ModifyChapters] Add ability for `--remove-chapters` to remove sections by timestamp
* [utils] Allow duration strings in `--match-filter`
* Add HDR information to formats
* Add negative option `--no-batch-file` by [Zirro](https://github.com/Zirro)
* Calculate more fields for merged formats
* Do not verify thumbnail URLs unless `--check-formats` is specified
* Don't create console for subprocesses on Windows
* Fix `--restrict-filename` when used with default template
* Fix `check_formats` output being written to stdout when `-qv`
* Fix bug in storyboards
* Fix conflict b/w id and ext in format selection
* Fix verbose head not showing custom configs
* Load archive only after printing verbose head
* Make `duration_string` and `resolution` available in --match-filter
* Re-implement deprecated option `--id`
* Reduce default `--socket-timeout`
* Write verbose header to logger
* [outtmpl] Fix bug in expanding environment variables
* [cookies] Local State should be opened as utf-8
* [extractor,utils] Detect more codecs/mimetypes
* [extractor] Detect `EXT-X-KEY` Apple FairPlay
* [utils] Use `importlib` to load plugins by [sulyi](https://github.com/sulyi)
* [http] Retry on socket timeout and show the last encountered error
* [fragment] Print error message when skipping fragment
* [aria2c] Fix `--skip-unavailable-fragment`
* [SponsorBlock] Obey `extractor-retries` and `sleep-requests`
* [Merger] Do not add `aac_adtstoasc` to non-hls audio
* [ModifyChapters] Do not mutate original chapters by [nihil-admirari](https://github.com/nihil-admirari)
* [devscripts/run_tests] Use markers to filter tests by [sulyi](https://github.com/sulyi)
* [7plus] Add cookie based authentication by [nyuszika7h](https://github.com/nyuszika7h)
* [AdobePass] Fix RCN MSO by [jfogelman](https://github.com/jfogelman)
* [CBC] Fix Gem livestream by [makeworld-the-better-one](https://github.com/makeworld-the-better-one)
* [CBC] Support CBC Gem member content by [makeworld-the-better-one](https://github.com/makeworld-the-better-one)
* [crunchyroll] Add season to flat-playlist
* [crunchyroll] Add support for `beta.crunchyroll` URLs and fix series URLs with language code
* [EUScreen] Add Extractor by [Ashish0804](https://github.com/Ashish0804)
* [Gronkh] Add extractor by [Ashish0804](https://github.com/Ashish0804)
* [hidive] Fix typo
* [Hotstar] Mention Dynamic Range in `format_id` by [Ashish0804](https://github.com/Ashish0804)
* [Hotstar] Raise appropriate error for DRM
* [instagram] Add login by [u-spec-png](https://github.com/u-spec-png)
* [instagram] Show appropriate error when login is needed
* [microsoftstream] Add extractor by [damianoamatruda](https://github.com/damianoamatruda), [nixklai](https://github.com/nixklai)
* [on24] Add extractor by [damianoamatruda](https://github.com/damianoamatruda)
* [patreon] Fix vimeo player regex by [zenerdi0de](https://github.com/zenerdi0de)
* [SkyNewsAU] Add extractor by [Ashish0804](https://github.com/Ashish0804)
* [tagesschau] Fix extractor by [u-spec-png](https://github.com/u-spec-png)
* [tbs] Add tbs live streams by [llacb47](https://github.com/llacb47)
* [tiktok] Fix typo and update tests
* [trovo] Support channel clips and VODs by [Ashish0804](https://github.com/Ashish0804)
* [Viafree] Add support for Finland by [18928172992817182](https://github.com/18928172992817182)
* [vimeo] Fix embedded `player.vimeo`
* [vlive:channel] Fix extraction by [kikuyan](https://github.com/kikuyan), [pukkandan](https://github.com/pukkandan)
* [youtube] Add auto-translated subtitles
* [youtube] Expose different formats with same itag
* [youtube:comments] Fix for new layout by [coletdjnz](https://github.com/coletdjnz)
* [cleanup] Cleanup bilibili code by [pukkandan](https://github.com/pukkandan), [u-spec-png](https://github.com/u-spec-png)
* [cleanup] Remove broken youtube login code
* [cleanup] Standardize timestamp formatting code
* [cleanup] Generalize `getcomments` implementation for extractors
* [cleanup] Simplify search extractors code
* [cleanup] misc


### 2021.10.10

* [downloader/ffmpeg] Fix bug in initializing `FFmpegPostProcessor`
* [minicurses] Fix when printing to file
* [downloader] Fix throttledratelimit
* [francetv] Fix extractor by [fstirlitz](https://github.com/fstirlitz), [sarnoud](https://github.com/sarnoud)
* [NovaPlay] Add extractor by [Bojidarist](https://github.com/Bojidarist)
* [ffmpeg] Revert "Set max probesize" - No longer needed
* [docs] Remove incorrect dependency on VC++10
* [build] Allow to release without changelog

### 2021.10.09

* Improved progress reporting
    * Separate `--console-title` and `--no-progress`
    * Add option `--progress` to show progress-bar even in quiet mode
    * Fix and refactor `minicurses` and use it for all progress reporting
    * Standardize use of terminal sequences and enable color support for windows 10
    * Add option `--progress-template` to customize progress-bar and console-title
    * Add postprocessor hooks and progress reporting
* [postprocessor] Add plugin support with option `--use-postprocessor`
* [extractor] Extract storyboards from SMIL manifests by [fstirlitz](https://github.com/fstirlitz)
* [outtmpl] Alternate form of format type `l` for `\n` delimited list
* [outtmpl] Format type `U` for unicode normalization
* [outtmpl] Allow empty output template to skip a type of file
* Merge webm formats into mkv if thumbnails are to be embedded
* [adobepass] Add RCN as MSO by [jfogelman](https://github.com/jfogelman)
* [ciscowebex] Add extractor by [damianoamatruda](https://github.com/damianoamatruda)
* [Gettr] Add extractor by [i6t](https://github.com/i6t)
* [GoPro] Add extractor by [i6t](https://github.com/i6t)
* [N1] Add extractor by [u-spec-png](https://github.com/u-spec-png)
* [Theta] Add video extractor by [alerikaisattera](https://github.com/alerikaisattera)
* [Veo] Add extractor by [i6t](https://github.com/i6t)
* [Vupload] Add extractor by [u-spec-png](https://github.com/u-spec-png)
* [bbc] Extract better quality videos by [ajj8](https://github.com/ajj8)
* [Bilibili] Add subtitle converter by [u-spec-png](https://github.com/u-spec-png)
* [CBC] Cleanup tests by [makeworld-the-better-one](https://github.com/makeworld-the-better-one)
* [Douyin] Rewrite extractor by [MinePlayersPE](https://github.com/MinePlayersPE)
* [Funimation] Fix for /v/ urls by [pukkandan](https://github.com/pukkandan), [Jules-A](https://github.com/Jules-A)
* [Funimation] Sort formats according to the relevant extractor-args
* [Hidive] Fix duplicate and incorrect formats
* [HotStarSeries] Fix cookies by [Ashish0804](https://github.com/Ashish0804)
* [LinkedInLearning] Add subtitles by [Ashish0804](https://github.com/Ashish0804)
* [Mediaite] Relax valid url by [coletdjnz](https://github.com/coletdjnz)
* [Newgrounds] Add age_limit and fix duration by [u-spec-png](https://github.com/u-spec-png)
* [Newgrounds] Fix view count on songs by [u-spec-png](https://github.com/u-spec-png)
* [parliamentlive.tv] Fix extractor by [u-spec-png](https://github.com/u-spec-png)
* [PolskieRadio] Fix extractors by [jakubadamw](https://github.com/jakubadamw), [u-spec-png](https://github.com/u-spec-png)
* [reddit] Add embedded url by [u-spec-png](https://github.com/u-spec-png)
* [reddit] Fix 429 by generating a random `reddit_session` by [AjaxGb](https://github.com/AjaxGb)
* [Rumble] Add RumbleChannelIE by [Ashish0804](https://github.com/Ashish0804)
* [soundcloud:playlist] Detect last page correctly
* [SovietsCloset] Add duration from m3u8 by [ChillingPepper](https://github.com/ChillingPepper)
* [Streamable] Add codecs by [u-spec-png](https://github.com/u-spec-png)
* [vidme] Remove extractor by [alerikaisattera](https://github.com/alerikaisattera)
* [youtube:tab] Fallback to API when webpage fails to download by [coletdjnz](https://github.com/coletdjnz)
* [youtube] Fix non-fatal errors in fetching player
* Fix `--flat-playlist` when neither IE nor id is known
* Fix `-f mp4` behaving differently from youtube-dl
* Workaround for bug in `ssl.SSLContext.load_default_certs`
* [aes] Improve performance slightly by [sulyi](https://github.com/sulyi)
* [cookies] Fix keyring fallback by [mbway](https://github.com/mbway)
* [embedsubtitle] Fix error when duration is unknown
* [ffmpeg] Fix error when subtitle file is missing
* [ffmpeg] Set max probesize to workaround AAC HLS stream issues by [shirt](https://github.com/shirt-dev)
* [FixupM3u8] Remove redundant run if merged is needed
* [hls] Fix decryption issues by [shirt](https://github.com/shirt-dev), [pukkandan](https://github.com/pukkandan)
* [http] Respect user-provided chunk size over extractor's
* [utils] Let traverse_obj accept functions as keys
* [docs] Add note about our custom ffmpeg builds
* [docs] Write embedding and contributing documentation by [pukkandan](https://github.com/pukkandan), [timethrow](https://github.com/timethrow)
* [update] Check for new version even if not updateable
* [build] Add more files to the tarball
* [build] Allow building with py2exe (and misc fixes)
* [build] Use pycryptodomex by [shirt](https://github.com/shirt-dev), [pukkandan](https://github.com/pukkandan)
* [cleanup] Some minor refactoring, improve docs and misc cleanup


### 2021.09.25

* Add new option `--netrc-location`
* [outtmpl] Allow alternate fields using `,`
* [outtmpl] Add format type `B` to treat the value as bytes, e.g. to limit the filename to a certain number of bytes
* Separate the options `--ignore-errors` and `--no-abort-on-error`
* Basic framework for simultaneous download of multiple formats by [nao20010128nao](https://github.com/nao20010128nao)
* [17live] Add 17.live extractor by [nao20010128nao](https://github.com/nao20010128nao)
* [bilibili] Add BiliIntlIE and BiliIntlSeriesIE by [Ashish0804](https://github.com/Ashish0804)
* [CAM4] Add extractor by [alerikaisattera](https://github.com/alerikaisattera)
* [Chingari] Add extractors by [Ashish0804](https://github.com/Ashish0804)
* [CGTN] Add extractor by [chao813](https://github.com/chao813)
* [damtomo] Add extractor by [nao20010128nao](https://github.com/nao20010128nao)
* [gotostage] Add extractor by [poschi3](https://github.com/poschi3)
* [Koo] Add extractor by [Ashish0804](https://github.com/Ashish0804)
* [Mediaite] Add Extractor by [Ashish0804](https://github.com/Ashish0804)
* [Mediaklikk] Add Extractor by [tmarki](https://github.com/tmarki), [mrx23dot](https://github.com/mrx23dot), [coletdjnz](https://github.com/coletdjnz)
* [MuseScore] Add Extractor by [Ashish0804](https://github.com/Ashish0804)
* [Newgrounds] Add NewgroundsUserIE and improve extractor by [u-spec-png](https://github.com/u-spec-png)
* [nzherald] Add NZHeraldIE by [coletdjnz](https://github.com/coletdjnz)
* [Olympics] Add replay extractor by [Ashish0804](https://github.com/Ashish0804)
* [Peertube] Add channel and playlist extractors by [u-spec-png](https://github.com/u-spec-png)
* [radlive] Add extractor by [nyuszika7h](https://github.com/nyuszika7h)
* [SovietsCloset] Add extractor by [ChillingPepper](https://github.com/ChillingPepper)
* [Streamanity] Add Extractor by [alerikaisattera](https://github.com/alerikaisattera)
* [Theta] Add extractor by [alerikaisattera](https://github.com/alerikaisattera)
* [Yandex] Add ZenYandexIE and ZenYandexChannelIE by [Ashish0804](https://github.com/Ashish0804)
* [9Now] handle episodes of series by [dalanmiller](https://github.com/dalanmiller)
* [AnimalPlanet] Fix extractor by [Sipherdrakon](https://github.com/Sipherdrakon)
* [Arte] Improve description extraction by [renalid](https://github.com/renalid)
* [atv.at] Use jwt for API by [NeroBurner](https://github.com/NeroBurner)
* [brightcove] Extract subtitles from manifests
* [CBC] Fix CBC Gem extractors by [makeworld-the-better-one](https://github.com/makeworld-the-better-one)
* [cbs] Report appropriate error for DRM
* [comedycentral] Support `collection-playlist` by [nixxo](https://github.com/nixxo)
* [DIYNetwork] Support new format by [Sipherdrakon](https://github.com/Sipherdrakon)
* [downloader/niconico] Pass custom headers by [nao20010128nao](https://github.com/nao20010128nao)
* [dw] Fix extractor
* [Fancode] Fix live streams by [zenerdi0de](https://github.com/zenerdi0de)
* [funimation] Fix for locations outside US by [Jules-A](https://github.com/Jules-A), [pukkandan](https://github.com/pukkandan)
* [globo] Fix GloboIE by [Ashish0804](https://github.com/Ashish0804)
* [HiDive] Fix extractor by [Ashish0804](https://github.com/Ashish0804)
* [Hotstar] Add referer for subs by [Ashish0804](https://github.com/Ashish0804)
* [itv] Fix extractor, add subtitles and thumbnails by [coletdjnz](https://github.com/coletdjnz), [sleaux-meaux](https://github.com/sleaux-meaux), [Vangelis66](https://github.com/Vangelis66)
* [lbry] Show error message from API response
* [Mxplayer] Use mobile API by [Ashish0804](https://github.com/Ashish0804)
* [NDR] Rewrite NDRIE by [Ashish0804](https://github.com/Ashish0804)
* [Nuvid] Fix extractor by [u-spec-png](https://github.com/u-spec-png)
* [Oreilly] Handle new web url by [MKSherbini](https://github.com/MKSherbini)
* [pbs] Fix subtitle extraction by [coletdjnz](https://github.com/coletdjnz), [gesa](https://github.com/gesa), [raphaeldore](https://github.com/raphaeldore)
* [peertube] Update instances by [u-spec-png](https://github.com/u-spec-png)
* [plutotv] Fix extractor for URLs with `/en`
* [reddit] Workaround for 429 by redirecting to old.reddit.com
* [redtube] Fix exts
* [soundcloud] Make playlist extraction lazy
* [soundcloud] Retry playlist pages on `502` error and update `_CLIENT_ID`
* [southpark] Fix SouthParkDE by [coletdjnz](https://github.com/coletdjnz)
* [SovietsCloset] Fix playlists for games with only named categories by [ConquerorDopy](https://github.com/ConquerorDopy)
* [SpankBang] Fix uploader by [f4pp3rk1ng](https://github.com/f4pp3rk1ng), [coletdjnz](https://github.com/coletdjnz)
* [tiktok] Use API to fetch higher quality video by [MinePlayersPE](https://github.com/MinePlayersPE), [llacb47](https://github.com/llacb47)
* [TikTokUser] Fix extractor using mobile API by [MinePlayersPE](https://github.com/MinePlayersPE), [llacb47](https://github.com/llacb47)
* [videa] Fix some extraction errors by [nyuszika7h](https://github.com/nyuszika7h)
* [VrtNU] Handle login errors by [llacb47](https://github.com/llacb47)
* [vrv] Don't raise error when thumbnails are missing
* [youtube] Cleanup authentication code by [coletdjnz](https://github.com/coletdjnz)
* [youtube] Fix `--mark-watched` with `--cookies-from-browser`
* [youtube] Improvements to JS player extraction and add extractor-args to skip it by [coletdjnz](https://github.com/coletdjnz)
* [youtube] Retry on 'Unknown Error' by [coletdjnz](https://github.com/coletdjnz)
* [youtube] Return full URL instead of just ID
* [youtube] Warn when trying to download clips
* [zdf] Improve format sorting
* [zype] Extract subtitles from the m3u8 manifest by [fstirlitz](https://github.com/fstirlitz)
* Allow `--force-write-archive` to work with `--flat-playlist`
* Download subtitles in order of `--sub-langs`
* Allow `0` in `--playlist-items`
* Handle more playlist errors with `-i`
* Fix `--no-get-comments`
* Fix `extra_info` being reused across runs
* Fix compat options `no-direct-merge` and `playlist-index`
* Dump files should obey `--trim-filename` by [sulyi](https://github.com/sulyi)
* [aes] Add `aes_gcm_decrypt_and_verify` by [sulyi](https://github.com/sulyi), [pukkandan](https://github.com/pukkandan)
* [aria2c] Fix IV for some AES-128 streams by [shirt](https://github.com/shirt-dev)
* [compat] Don't ignore `HOME` (if set) on windows
* [cookies] Make browser names case insensitive
* [cookies] Print warning for cookie decoding error only once
* [extractor] Fix root-relative URLs in MPD by [DigitalDJ](https://github.com/DigitalDJ)
* [ffmpeg] Add `aac_adtstoasc` when merging if needed
* [fragment,aria2c] Generalize and refactor some code
* [fragment] Avoid repeated request for AES key
* [fragment] Fix range header when using `-N` and media sequence by [shirt](https://github.com/shirt-dev)
* [hls,aes] Fallback to native implementation for AES-CBC and detect `Cryptodome` in addition to `Crypto`
* [hls] Byterange + AES128 is supported by native downloader
* [ModifyChapters] Improve sponsor chapter merge algorithm by [nihil-admirari](https://github.com/nihil-admirari)
* [ModifyChapters] Minor fixes
* [WebVTT] Adjust parser to accommodate PBS subtitles
* [utils] Improve `extract_timezone` by [dirkf](https://github.com/dirkf)
* [options] Fix `--no-config` and refactor reading of config files
* [options] Strip spaces and ignore empty entries in list-like switches
* [test/cookies] Improve logging
* [build] Automate more of the release process by [animelover1984](https://github.com/animelover1984), [pukkandan](https://github.com/pukkandan)
* [build] Fix sha256 by [nihil-admirari](https://github.com/nihil-admirari)
* [build] Bring back brew taps by [nao20010128nao](https://github.com/nao20010128nao)
* [build] Provide `--onedir` zip for windows
* [cleanup,docs] Add deprecation warning in docs for some counter intuitive behaviour
* [cleanup] Fix line endings for `nebula.py` by [glenn-slayden](https://github.com/glenn-slayden)
* [cleanup] Improve `make clean-test` by [sulyi](https://github.com/sulyi)
* [cleanup] Misc


### 2021.09.02

* **Native SponsorBlock** implementation by [nihil-admirari](https://github.com/nihil-admirari), [pukkandan](https://github.com/pukkandan)
    * `--sponsorblock-remove CATS` removes specified chapters from file
    * `--sponsorblock-mark CATS` marks the specified sponsor sections as chapters
    * `--sponsorblock-chapter-title TMPL` to specify sponsor chapter template
    * `--sponsorblock-api URL` to use a different API
    * No re-encoding is done unless `--force-keyframes-at-cuts` is used
    * The fetched sponsor sections are written to the infojson
    * Deprecates: `--sponskrub`, `--no-sponskrub`, `--sponskrub-cut`, `--no-sponskrub-cut`, `--sponskrub-force`, `--no-sponskrub-force`, `--sponskrub-location`, `--sponskrub-args`
* Split `--embed-chapters` from `--embed-metadata` (it still implies the former by default)
* Add option `--remove-chapters` to remove arbitrary chapters by [nihil-admirari](https://github.com/nihil-admirari), [pukkandan](https://github.com/pukkandan)
* Add option `--force-keyframes-at-cuts` for more accurate cuts when removing and splitting chapters by [nihil-admirari](https://github.com/nihil-admirari)
* Let `--match-filter` reject entries early
    * Makes redundant: `--match-title`, `--reject-title`, `--min-views`, `--max-views`
* [lazy_extractor] Improvements (It now passes all tests)
    * Bugfix for when plugin directory doesn't exist by [kidonng](https://github.com/kidonng)
    * Create instance only after pre-checking archive
    * Import actual class if an attribute is accessed
    * Fix `suitable` and add flake8 test
* [downloader/ffmpeg] Experimental support for DASH manifests (including live)
    * Your ffmpeg must have [this patch](https://github.com/FFmpeg/FFmpeg/commit/3249c757aed678780e22e99a1a49f4672851bca9) applied for YouTube DASH to work
* [downloader/ffmpeg] Allow passing custom arguments before `-i`
* [BannedVideo] Add extractor by [smege1001](https://github.com/smege1001), [blackjack4494](https://github.com/blackjack4494), [pukkandan](https://github.com/pukkandan)
* [bilibili] Add category extractor by [animelover1984](https://github.com/animelover1984)
* [Epicon] Add extractors by [Ashish0804](https://github.com/Ashish0804)
* [filmmodu] Add extractor by [mzbaulhaque](https://github.com/mzbaulhaque)
* [GabTV] Add extractor by [Ashish0804](https://github.com/Ashish0804)
* [Hungama] Fix `HungamaSongIE` and add `HungamaAlbumPlaylistIE` by [Ashish0804](https://github.com/Ashish0804)
* [ManotoTV] Add new extractors by [tandy1000](https://github.com/tandy1000)
* [Niconico] Add Search extractors by [animelover1984](https://github.com/animelover1984), [pukkandan](https://github.com/pukkandan)
* [Patreon] Add `PatreonUserIE` by [zenerdi0de](https://github.com/zenerdi0de)
* [peloton] Add extractor by [IONECarter](https://github.com/IONECarter), [capntrips](https://github.com/capntrips), [pukkandan](https://github.com/pukkandan)
* [ProjectVeritas] Add extractor by [Ashish0804](https://github.com/Ashish0804)
* [radiko] Add extractors by [nao20010128nao](https://github.com/nao20010128nao)
* [StarTV] Add extractor for `startv.com.tr` by [mrfade](https://github.com/mrfade), [coletdjnz](https://github.com/coletdjnz)
* [tiktok] Add `TikTokUserIE` by [Ashish0804](https://github.com/Ashish0804), [pukkandan](https://github.com/pukkandan)
* [Tokentube] Add extractor by [u-spec-png](https://github.com/u-spec-png)
* [TV2Hu] Fix `TV2HuIE` and add `TV2HuSeriesIE` by [Ashish0804](https://github.com/Ashish0804)
* [voicy] Add extractor by [nao20010128nao](https://github.com/nao20010128nao)
* [adobepass] Fix Verizon SAML login by [nyuszika7h](https://github.com/nyuszika7h), [ParadoxGBB](https://github.com/ParadoxGBB)
* [afreecatv] Fix adult VODs by [wlritchi](https://github.com/wlritchi)
* [afreecatv] Tolerate failure to parse date string by [wlritchi](https://github.com/wlritchi)
* [aljazeera] Fix extractor by [MinePlayersPE](https://github.com/MinePlayersPE)
* [ATV.at] Fix extractor for ATV.at by [NeroBurner](https://github.com/NeroBurner), [coletdjnz](https://github.com/coletdjnz)
* [bitchute] Fix test by [mahanstreamer](https://github.com/mahanstreamer)
* [camtube] Remove obsolete extractor by [alerikaisattera](https://github.com/alerikaisattera)
* [CDA] Add more formats by [u-spec-png](https://github.com/u-spec-png)
* [eroprofile] Fix page skipping in albums by [jhwgh1968](https://github.com/jhwgh1968)
* [facebook] Fix format sorting
* [facebook] Fix metadata extraction by [kikuyan](https://github.com/kikuyan)
* [facebook] Update onion URL by [Derkades](https://github.com/Derkades)
* [HearThisAtIE] Fix extractor by [Ashish0804](https://github.com/Ashish0804)
* [instagram] Add referrer to prevent throttling by [u-spec-png](https://github.com/u-spec-png), [kikuyan](https://github.com/kikuyan)
* [iwara.tv] Extract more metadata by [BunnyHelp](https://github.com/BunnyHelp)
* [iwara] Add thumbnail by [i6t](https://github.com/i6t)
* [kakao] Fix extractor
* [mediaset] Fix extraction for some videos by [nyuszika7h](https://github.com/nyuszika7h)
* [Motherless] Fix extractor by [coletdjnz](https://github.com/coletdjnz)
* [Nova] fix extractor by [std-move](https://github.com/std-move)
* [ParamountPlus] Fix geo verification by [shirt](https://github.com/shirt-dev)
* [peertube] handle new video URL format by [Chocobozzz](https://github.com/Chocobozzz)
* [pornhub] Separate and fix playlist extractor by [mzbaulhaque](https://github.com/mzbaulhaque)
* [reddit] Fix for quarantined subreddits by [ouwou](https://github.com/ouwou)
* [ShemarooMe] Fix extractor by [Ashish0804](https://github.com/Ashish0804)
* [soundcloud] Refetch `client_id` on 403
* [tiktok] Fix metadata extraction
* [TV2] Fix extractor by [Ashish0804](https://github.com/Ashish0804)
* [tv5mondeplus] Fix extractor by [korli](https://github.com/korli)
* [VH1,TVLand] Fix extractors by [Sipherdrakon](https://github.com/Sipherdrakon)
* [Viafree] Fix extractor and extract subtitles by [coletdjnz](https://github.com/coletdjnz)
* [XHamster] Extract `uploader_id` by [octotherp](https://github.com/octotherp)
* [youtube] Add `shorts` to `_VALID_URL`
* [youtube] Add av01 itags to known formats list by [blackjack4494](https://github.com/blackjack4494)
* [youtube] Extract error messages from HTTPError response by [coletdjnz](https://github.com/coletdjnz)
* [youtube] Fix subtitle names
* [youtube] Prefer audio stream that YouTube considers default
* [youtube] Remove annotations and deprecate `--write-annotations` by [coletdjnz](https://github.com/coletdjnz)
* [Zee5] Fix extractor and add subtitles by [Ashish0804](https://github.com/Ashish0804)
* [aria2c] Obey `--rate-limit`
* [EmbedSubtitle] Continue even if some files are missing
* [extractor] Better error message for DRM
* [extractor] Common function `_match_valid_url`
* [extractor] Show video id in error messages if possible
* [FormatSort] Remove priority of `lang`
* [options] Add `_set_from_options_callback`
* [SubtitleConvertor] Fix bug during subtitle conversion
* [utils] Add `parse_qs`
* [webvtt] Fix timestamp overflow adjustment by [fstirlitz](https://github.com/fstirlitz)
* Bugfix for `--replace-in-metadata`
* Don't try to merge with final extension
* Fix `--force-overwrites` when using `-k`
* Fix `--no-prefer-free-formats` by [CeruleanSky](https://github.com/CeruleanSky)
* Fix `-F` for extractors that directly return url
* Fix `-J` when there are failed videos
* Fix `extra_info` being reused across runs
* Fix `playlist_index` not obeying `playlist_start` and add tests
* Fix resuming of single formats when using `--no-part`
* Revert erroneous use of the `Content-Length` header by [fstirlitz](https://github.com/fstirlitz)
* Use `os.replace` where applicable by; paulwrubel
* [build] Add homebrew taps `yt-dlp/taps/yt-dlp` by [nao20010128nao](https://github.com/nao20010128nao)
* [build] Fix bug in making `yt-dlp.tar.gz`
* [docs] Fix some typos by [pukkandan](https://github.com/pukkandan), [zootedb0t](https://github.com/zootedb0t)
* [cleanup] Replace improper use of tab in trovo by [glenn-slayden](https://github.com/glenn-slayden)


### 2021.08.10

* Add option `--replace-in-metadata`
* Add option `--no-simulate` to not simulate even when `--print` or `--list...` are used - Deprecates `--print-json`
* Allow entire infodict to be printed using `%()s` - makes `--dump-json` redundant
* Allow multiple `--exec` and `--exec-before-download`
* Add regex to `--match-filter`
* Add all format filtering operators also to `--match-filter` by [max-te](https://github.com/max-te)
* Add compat-option `no-keep-subs`
* [adobepass] Add MSO Cablevision by [Jessecar96](https://github.com/Jessecar96)
* [BandCamp] Add BandcampMusicIE by [Ashish0804](https://github.com/Ashish0804)
* [blackboardcollaborate] Add new extractor by [mzbaulhaque](https://github.com/mzbaulhaque)
* [eroprofile] Add album downloader by [jhwgh1968](https://github.com/jhwgh1968)
* [mirrativ] Add extractors by [nao20010128nao](https://github.com/nao20010128nao)
* [openrec] Add extractors by [nao20010128nao](https://github.com/nao20010128nao)
* [nbcolympics:stream] Fix extractor by [nchilada](https://github.com/nchilada), [pukkandan](https://github.com/pukkandan)
* [nbcolympics] Update extractor for 2020 olympics by [wesnm](https://github.com/wesnm)
* [paramountplus] Separate extractor and fix some titles by [shirt](https://github.com/shirt-dev), [pukkandan](https://github.com/pukkandan)
* [RCTIPlus] Support events and TV by [MinePlayersPE](https://github.com/MinePlayersPE)
* [Newgrounds] Improve extractor and fix playlist by [u-spec-png](https://github.com/u-spec-png)
* [aenetworks] Update `_THEPLATFORM_KEY` and `_THEPLATFORM_SECRET` by [wesnm](https://github.com/wesnm)
* [crunchyroll] Fix thumbnail by [funniray](https://github.com/funniray)
* [HotStar] Use API for metadata and extract subtitles by [Ashish0804](https://github.com/Ashish0804)
* [instagram] Fix comments extraction by [u-spec-png](https://github.com/u-spec-png)
* [peertube] Fix videos without description by [u-spec-png](https://github.com/u-spec-png)
* [twitch:clips] Extract `display_id` by [dirkf](https://github.com/dirkf)
* [viki] Print error message from API request
* [Vine] Remove invalid formats by [u-spec-png](https://github.com/u-spec-png)
* [VrtNU] Fix XSRF token by [pgaig](https://github.com/pgaig)
* [vrv] Fix thumbnail extraction by [funniray](https://github.com/funniray)
* [youtube] Add extractor-arg `include-live-dash` to show live dash formats
* [youtube] Improve signature function detection by [PSlava](https://github.com/PSlava)
* [youtube] Raise appropriate error when API pages can't be downloaded
* Ensure `_write_ytdl_file` closes file handle on error
* Fix `--compat-options filename` by [stdedos](https://github.com/stdedos)
* Fix issues with infodict sanitization
* Fix resuming when using `--no-part`
* Fix wrong extension for intermediate files
* Handle `BrokenPipeError` by [kikuyan](https://github.com/kikuyan)
* Show libraries present in verbose head
* [extractor] Detect `sttp` as subtitles in MPD by [fstirlitz](https://github.com/fstirlitz)
* [extractor] Reset non-repeating warnings per video
* [ffmpeg] Fix streaming `mp4` to `stdout`
* [ffpmeg] Allow `--ffmpeg-location` to be a file with different name
* [utils] Fix `InAdvancePagedList.__getitem__`
* [utils] Fix `traverse_obj` depth when `is_user_input`
* [webvtt] Merge daisy-chained duplicate cues by [fstirlitz](https://github.com/fstirlitz)
* [build] Use custom build of `pyinstaller` by [shirt](https://github.com/shirt-dev)
* [tests:download] Add batch testing for extractors (`test_YourExtractor_all`)
* [docs] Document which fields `--add-metadata` adds to the file
* [docs] Fix some mistakes and improve doc
* [cleanup] Misc code cleanup


### 2021.08.02

* Add logo, banner and donate links
* [outtmpl] Expand and escape environment variables
* [outtmpl] Add format types `j` (json), `l` (comma delimited list), `q` (quoted for terminal)
* [downloader] Allow streaming some unmerged formats to stdout using ffmpeg
* [youtube] **Age-gate bypass**
    * Add `agegate` clients by [pukkandan](https://github.com/pukkandan), [MinePlayersPE](https://github.com/MinePlayersPE)
    * Add `thirdParty` to agegate clients to bypass more videos
    * Simplify client definitions, expose `embedded` clients
    * Improve age-gate detection by [coletdjnz](https://github.com/coletdjnz)
    * Fix default global API key by [coletdjnz](https://github.com/coletdjnz)
    * Add `creator` clients for age-gate bypass using unverified accounts by [zerodytrash](https://github.com/zerodytrash), [coletdjnz](https://github.com/coletdjnz), [pukkandan](https://github.com/pukkandan)
* [adobepass] Add MSO Sling TV by [wesnm](https://github.com/wesnm)
* [CBS] Add ParamountPlusSeriesIE by [Ashish0804](https://github.com/Ashish0804)
* [dplay] Add `ScienceChannelIE` by [Sipherdrakon](https://github.com/Sipherdrakon)
* [UtreonIE] Add extractor by [Ashish0804](https://github.com/Ashish0804)
* [youtube] Add `mweb` client by [coletdjnz](https://github.com/coletdjnz)
* [youtube] Add `player_client=all`
* [youtube] Force `hl=en` for comments by [coletdjnz](https://github.com/coletdjnz)
* [youtube] Fix format sorting when using alternate clients
* [youtube] Misc cleanup by [pukkandan](https://github.com/pukkandan), [coletdjnz](https://github.com/coletdjnz)
* [youtube] Extract SAPISID only once
* [CBS] Add fallback by [llacb47](https://github.com/llacb47), [pukkandan](https://github.com/pukkandan)
* [Hotstar] Support cookies by [Ashish0804](https://github.com/Ashish0804)
* [HotStarSeriesIE] Fix regex by [Ashish0804](https://github.com/Ashish0804)
* [bilibili] Improve `_VALID_URL`
* [mediaset] Fix extraction by [nixxo](https://github.com/nixxo)
* [Mxplayer] Add h265 formats by [Ashish0804](https://github.com/Ashish0804)
* [RCTIPlus] Remove PhantomJS dependency by [MinePlayersPE](https://github.com/MinePlayersPE)
* [tenplay] Add MA15+ age limit by [pento](https://github.com/pento)
* [vidio] Fix login error detection by [MinePlayersPE](https://github.com/MinePlayersPE)
* [vimeo] Better extraction of original file by [Ashish0804](https://github.com/Ashish0804)
* [generic] Support KVS player (replaces ThisVidIE) by [rigstot](https://github.com/rigstot)
* Add compat-option `no-clean-infojson`
* Remove `asr` appearing twice in `-F`
* Set `home:` as the default key for `-P`
* [utils] Fix slicing of reversed `LazyList`
* [FormatSort] Fix bug for audio with unknown codec
* [test:download] Support testing with `ignore_no_formats_error`
* [cleanup] Refactor some code


### 2021.07.24

* [youtube:tab] Extract video duration early
* [downloader] Pass `info_dict` to `progress_hook`s
* [youtube] Fix age-gated videos for API clients when cookies are supplied by [coletdjnz](https://github.com/coletdjnz)
* [youtube] Disable `get_video_info` age-gate workaround - This endpoint seems to be completely dead
* [youtube] Try all clients even if age-gated
* [youtube] Fix subtitles only being extracted from the first client
* [youtube] Simplify `_get_text`
* [cookies] bugfix for microsoft edge on macOS
* [cookies] Handle `sqlite` `ImportError` gracefully by [mbway](https://github.com/mbway)
* [cookies] Handle errors when importing `keyring`

### 2021.07.21

* **Add option `--cookies-from-browser`** to load cookies from a browser by [mbway](https://github.com/mbway)
    * Usage: `--cookies-from-browser BROWSER[:PROFILE_NAME_OR_PATH]`
    * Also added `--no-cookies-from-browser`
    * To decrypt chromium cookies, `keyring` is needed for UNIX and `pycryptodome` for Windows
* Add option `--exec-before-download`
* Add field `live_status`
* [FFmpegMetadata] Add language of each stream and some refactoring
* [douyin] Add extractor by [pukkandan](https://github.com/pukkandan), [pyx](https://github.com/pyx)
* [pornflip] Add extractor by [mzbaulhaque](https://github.com/mzbaulhaque)
* **[youtube] Extract data from multiple clients** by [pukkandan](https://github.com/pukkandan), [coletdjnz](https://github.com/coletdjnz)
    * `player_client` now accepts multiple clients
    * Default `player_client` = `android,web`
        * This uses twice as many requests, but avoids throttling for most videos while also not losing any formats
    * Music clients can be specifically requested and is enabled by default if `music.youtube.com`
    * Added `player_client=ios` (Known issue: formats from ios are not sorted correctly)
    * Add age-gate bypass for android and ios clients
* [youtube] Extract more thumbnails
    * The thumbnail URLs are hard-coded and their actual existence is tested lazily
    * Added option `--no-check-formats` to not test them
* [youtube] Misc fixes
    * Improve extraction of livestream metadata by [pukkandan](https://github.com/pukkandan), [krichbanana](https://github.com/krichbanana)
    * Hide live dash formats since they can't be downloaded anyway
    * Fix authentication when using multiple accounts by [coletdjnz](https://github.com/coletdjnz)
    * Fix controversial videos when requested via API by [coletdjnz](https://github.com/coletdjnz)
    * Fix session index extraction and headers for non-web player clients by [coletdjnz](https://github.com/coletdjnz)
    * Make `--extractor-retries` work for more errors
    * Fix sorting of 3gp format
    * Sanity check `chapters` (and refactor related code)
    * Make `parse_time_text` and `_extract_chapters` non-fatal
    * Misc cleanup and bug fixes by [coletdjnz](https://github.com/coletdjnz)
* [youtube:tab] Fix channels tab
* [youtube:tab] Extract playlist availability by [coletdjnz](https://github.com/coletdjnz)
* **[youtube:comments] Move comment extraction to new API** by [coletdjnz](https://github.com/coletdjnz)
    * Adds extractor-args `comment_sort` (`top`/`new`), `max_comments`, `max_comment_depth`
* [youtube:comments] Fix `is_favorited`, improve `like_count` parsing by [coletdjnz](https://github.com/coletdjnz)
* [BravoTV] Improve metadata extraction by [kevinoconnor7](https://github.com/kevinoconnor7)
* [crunchyroll:playlist] Force http
* [yahoo:gyao:player] Relax `_VALID_URL` by [nao20010128nao](https://github.com/nao20010128nao)
* [nebula] Authentication via tokens from cookie jar by [hheimbuerger](https://github.com/hheimbuerger), [TpmKranz](https://github.com/TpmKranz)
* [RTP] Fix extraction and add subtitles by [fstirlitz](https://github.com/fstirlitz)
* [viki] Rewrite extractors and add extractor-arg `video_types` to `vikichannel` by [zackmark29](https://github.com/zackmark29), [pukkandan](https://github.com/pukkandan)
* [vlive] Extract thumbnail directly in addition to the one from Naver
* [generic] Extract previously missed subtitles by [fstirlitz](https://github.com/fstirlitz)
* [generic] Extract everything in the SMIL manifest and detect discarded subtitles by [fstirlitz](https://github.com/fstirlitz)
* [embedthumbnail] Fix `_get_thumbnail_resolution`
* [metadatafromfield] Do not detect numbers as field names
* Fix selectors `all`, `mergeall` and add tests
* Errors in playlist extraction should obey `--ignore-errors`
* Fix bug where `original_url` was not propagated when `_type`=`url`
* Revert "Merge webm formats into mkv if thumbnails are to be embedded (#173)"
    * This was wrongly checking for `write_thumbnail`
* Improve `extractor_args` parsing
* Rename `NOTE` in `-F` to `MORE INFO` since it's often confused to be the same as `format_note`
* Add `only_once` param for `write_debug` and `report_warning`
* [extractor] Allow extracting multiple groups in `_search_regex` by [fstirlitz](https://github.com/fstirlitz)
* [utils] Improve `traverse_obj`
* [utils] Add `variadic`
* [utils] Improve `js_to_json` comment regex by [fstirlitz](https://github.com/fstirlitz)
* [webtt] Fix timestamps
* [compat] Remove unnecessary code
* [docs] fix default of multistreams


### 2021.07.07

* Merge youtube-dl: Upto [commit/a803582](https://github.com/ytdl-org/youtube-dl/commit/a8035827177d6b59aca03bd717acb6a9bdd75ada)
* Add `--extractor-args` to pass some extractor-specific arguments. See [readme](https://github.com/yt-dlp/yt-dlp#extractor-arguments)
    * Add extractor option `skip` for `youtube`, e.g. `--extractor-args youtube:skip=hls,dash`
    * Deprecates `--youtube-skip-dash-manifest`, `--youtube-skip-hls-manifest`, `--youtube-include-dash-manifest`, `--youtube-include-hls-manifest`
* Allow `--list...` options to work with `--print`, `--quiet` and other `--list...` options
* [youtube] Use `player` API for additional video extraction requests by [coletdjnz](https://github.com/coletdjnz)
    * **Fixes youtube premium music** (format 141) extraction
    * Adds extractor option `player_client` = `web`/`android`
        * **`--extractor-args youtube:player_client=android` works around the throttling** for the time-being
    * Adds extractor option `player_skip=config`
    * Adds age-gate fallback using embedded client
* [youtube] Choose correct Live chat API for upcoming streams by [krichbanana](https://github.com/krichbanana)
* [youtube] Fix subtitle names for age-gated videos
* [youtube:comments] Fix error handling and add `itct` to params by [coletdjnz](https://github.com/coletdjnz)
* [youtube_live_chat] Fix download with cookies by [siikamiika](https://github.com/siikamiika)
* [youtube_live_chat] use `clickTrackingParams` by [siikamiika](https://github.com/siikamiika)
* [Funimation] Rewrite extractor
    * Add `FunimationShowIE` by [Mevious](https://github.com/Mevious)
    * **Treat the different versions of an episode as different formats of a single video**
        * This changes the video `id` and will break break existing archives
        * Compat option `seperate-video-versions` to fall back to old behavior including using the old video ids
    * Support direct `/player/` URL
    * Extractor options `language` and `version` to pre-select them during extraction
        * These options may be removed in the future if we can extract all formats without additional network requests
        * Do not rely on these for format selection and use `-f` filters instead
* [AdobePass] Add Spectrum MSO by [kevinoconnor7](https://github.com/kevinoconnor7), [ohmybahgosh](https://github.com/ohmybahgosh)
* [facebook] Extract description and fix title
* [fancode] Fix extraction, support live and allow login with refresh token by [zenerdi0de](https://github.com/zenerdi0de)
* [plutotv] Improve `_VALID_URL`
* [RCTIPlus] Add extractor by [MinePlayersPE](https://github.com/MinePlayersPE)
* [Soundcloud] Allow login using oauth token by [blackjack4494](https://github.com/blackjack4494)
* [TBS] Support livestreams by [llacb47](https://github.com/llacb47)
* [videa] Fix extraction by [nyuszika7h](https://github.com/nyuszika7h)
* [yahoo] Fix extraction by [llacb47](https://github.com/llacb47), [pukkandan](https://github.com/pukkandan)
* Process videos when using `--ignore-no-formats-error` by [krichbanana](https://github.com/krichbanana)
* Fix `--throttled-rate` when using `--load-info-json`
* Fix `--flat-playlist` when entry has no `ie_key`
* Fix `check_formats` catching `ExtractorError` instead of `DownloadError`
* Fix deprecated option `--list-formats-old`
* [downloader/ffmpeg] Fix `--ppa` when using simultaneous download
* [extractor] Prevent unnecessary download of hls manifests and refactor `hls_split_discontinuity`
* [fragment] Handle status of download and errors in threads correctly; and minor refactoring
* [thumbnailsconvertor] Treat `jpeg` as `jpg`
* [utils] Fix issues with `LazyList` reversal
* [extractor] Allow extractors to set their own login hint
* [cleanup] Simplify format selector code with `LazyList` and `yield from`
* [cleanup] Clean `extractor.common._merge_subtitles` signature
* [cleanup] Fix some typos


### 2021.06.23

* Merge youtube-dl: Upto [commit/379f52a](https://github.com/ytdl-org/youtube-dl/commit/379f52a4954013767219d25099cce9e0f9401961)
* **Add option `--throttled-rate`** below which video data is re-extracted
* [fragment] **Merge during download for `-N`**, and refactor `hls`/`dash`
* [websockets] Add `WebSocketFragmentFD` by [nao20010128nao](https://github.com/nao20010128nao), [pukkandan](https://github.com/pukkandan)
* Allow `images` formats in addition to video/audio
* [downloader/mhtml] Add new downloader for slideshows/storyboards by [fstirlitz](https://github.com/fstirlitz)
* [youtube] Temporary **fix for age-gate**
* [youtube] Support ongoing live chat by [siikamiika](https://github.com/siikamiika)
* [youtube] Improve SAPISID cookie handling by [coletdjnz](https://github.com/coletdjnz)
* [youtube] Login is not needed for `:ytrec`
* [youtube] Non-fatal alert reporting for unavailable videos page by [coletdjnz](https://github.com/coletdjnz)
* [twitcasting] Websocket support by [nao20010128nao](https://github.com/nao20010128nao)
* [mediasite] Extract slides by [fstirlitz](https://github.com/fstirlitz)
* [funimation] Extract subtitles
* [pornhub] Extract `cast`
* [hotstar] Use server time for authentication instead of local time
* [EmbedThumbnail] Fix for already downloaded thumbnail
* [EmbedThumbnail] Add compat-option `embed-thumbnail-atomicparsley`
* Expand `--check-formats` to thumbnails
* Fix id sanitization in filenames
* Skip fixup of existing files and add `--fixup force` to force it
* Better error handling of syntax errors in `-f`
* Use `NamedTemporaryFile` for `--check-formats`
* [aria2c] Lower `--min-split-size` for HTTP downloads
* [options] Rename `--add-metadata` to `--embed-metadata`
* [utils] Improve `LazyList` and add tests
* [build] Build Windows x86 version with py3.7 and remove redundant tests by [pukkandan](https://github.com/pukkandan), [shirt](https://github.com/shirt-dev)
* [docs] Clarify that `--embed-metadata` embeds chapter markers
* [cleanup] Refactor fixup


### 2021.06.09

* Fix bug where `%(field)d` in filename template throws error
* [outtmpl] Improve offset parsing
* [test] More rigorous tests for `prepare_filename`

### 2021.06.08

* Remove support for obsolete Python versions: Only 3.6+ is now supported
* Merge youtube-dl: Upto [commit/c2350ca](https://github.com/ytdl-org/youtube-dl/commit/c2350cac243ba1ec1586fe85b0d62d1b700047a2)
* [hls] Fix decryption for multithreaded downloader
* [extractor] Fix pre-checking archive for some extractors
* [extractor] Fix FourCC fallback when parsing ISM by [fstirlitz](https://github.com/fstirlitz)
* [twitcasting] Add TwitCastingUserIE, TwitCastingLiveIE by [pukkandan](https://github.com/pukkandan), [nao20010128nao](https://github.com/nao20010128nao)
* [vidio] Add VidioPremierIE and VidioLiveIE by [MinePlayersPE](Https://github.com/MinePlayersPE)
* [viki] Fix extraction from [ytdl-org/youtube-dl@59e583f](https://github.com/ytdl-org/youtube-dl/commit/59e583f7e8530ca92776c866897d895c072e2a82)
* [youtube] Support shorts URL
* [zoom] Extract transcripts as subtitles
* Add field `original_url` with the user-inputted URL
* Fix and refactor `prepare_outtmpl`
* Make more fields available for `--print` when used with `--flat-playlist`
* [utils] Generalize `traverse_dict` to `traverse_obj`
* [downloader/ffmpeg] Hide FFmpeg banner unless in verbose mode by [fstirlitz](https://github.com/fstirlitz)
* [build] Release `yt-dlp.tar.gz`
* [build,update] Add GNU-style SHA512 and prepare updater for similar SHA256 by [nihil-admirari](https://github.com/nihil-admirari)
* [pyinst] Show Python version in exe metadata by [nihil-admirari](https://github.com/nihil-admirari)
* [docs] Improve documentation of dependencies
* [cleanup] Mark unused files
* [cleanup] Point all shebang to `python3` by [fstirlitz](https://github.com/fstirlitz)
* [cleanup] Remove duplicate file `trovolive.py`


### 2021.06.01

* Merge youtube-dl: Upto [commit/d495292](https://github.com/ytdl-org/youtube-dl/commit/d495292852b6c2f1bd58bc2141ff2b0265c952cf)
* Pre-check archive and filters during playlist extraction
* Handle Basic Auth `user:pass` in URLs by [hhirtz](https://github.com/hhirtz) and [pukkandan](https://github.com/pukkandan)
* [archiveorg] Add YoutubeWebArchiveIE by [coletdjnz](https://github.com/coletdjnz) and [alex-gedeon](https://github.com/alex-gedeon)
* [fancode] Add extractor by [rhsmachine](https://github.com/rhsmachine)
* [patreon] Support vimeo embeds by [rhsmachine](https://github.com/rhsmachine)
* [Saitosan] Add new extractor by [llacb47](https://github.com/llacb47)
* [ShemarooMe] Add extractor by [Ashish0804](https://github.com/Ashish0804) and [pukkandan](https://github.com/pukkandan)
* [telemundo] Add extractor by [king-millez](https://github.com/king-millez)
* [SonyLIV] Add SonyLIVSeriesIE and subtitle support by [Ashish0804](https://github.com/Ashish0804)
* [Hotstar] Add HotStarSeriesIE by [Ashish0804](https://github.com/Ashish0804)
* [Voot] Add VootSeriesIE by [Ashish0804](https://github.com/Ashish0804)
* [vidio] Support login and premium videos by [MinePlayersPE](https://github.com/MinePlayersPE)
* [fragment] When using `-N`, do not keep the fragment content in memory
* [ffmpeg] Download and merge in a single step if possible
* [ThumbnailsConvertor] Support conversion to `png` and make it the default by [louie-github](https://github.com/louie-github)
* [VideoConvertor] Generalize with remuxer and allow conditional recoding
* [EmbedThumbnail] Embed in `mp4`/`m4a` using mutagen by [tripulse](https://github.com/tripulse) and [pukkandan](https://github.com/pukkandan)
* [EmbedThumbnail] Embed if any thumbnail was downloaded, not just the best
* [EmbedThumbnail] Correctly escape filename
* [update] replace self without launching a subprocess in windows
* [update] Block further update for unsupported systems
* Refactor `__process_playlist` by creating `LazyList`
* Write messages to `stderr` when both `quiet` and `verbose`
* Sanitize and sort playlist thumbnails
* Remove `None` values from `info.json`
* [extractor] Always prefer native hls downloader by default
* [extractor] Skip subtitles without URI in m3u8 manifests by [hheimbuerger](https://github.com/hheimbuerger)
* [extractor] Functions to parse `socket.io` response as `json` by [pukkandan](https://github.com/pukkandan) and [llacb47](https://github.com/llacb47)
* [extractor] Allow `note=False` when extracting manifests
* [utils] Escape URLs in `sanitized_Request`, not `sanitize_url`
* [hls] Disable external downloader for `webtt`
* [youtube] `/live` URLs should raise error if channel is not live
* [youtube] Bug fixes
* [zee5] Fix m3u8 formats' extension
* [ard] Allow URLs without `-` before id by [olifre](https://github.com/olifre)
* [cleanup] `YoutubeDL._match_entry`
* [cleanup] Refactor updater
* [cleanup] Refactor ffmpeg convertors
* [cleanup] setup.py


### 2021.05.20

* **Youtube improvements**:
    * Support youtube music `MP`, `VL` and `browse` pages
    * Extract more formats for youtube music by [craftingmod](https://github.com/craftingmod), [coletdjnz](https://github.com/coletdjnz) and [pukkandan](https://github.com/pukkandan)
    * Extract multiple subtitles in same language by [pukkandan](https://github.com/pukkandan) and [tpikonen](https://github.com/tpikonen)
    * Redirect channels that doesn't have a `videos` tab to their `UU` playlists
    * Support in-channel search
    * Sort audio-only formats correctly
    * Always extract `maxresdefault` thumbnail
    * Extract audio language
    * Add subtitle language names by [nixxo](https://github.com/nixxo) and [tpikonen](https://github.com/tpikonen)
    * Show alerts only from the final webpage
    * Add `html5=1` param to `get_video_info` page requests by [coletdjnz](https://github.com/coletdjnz)
    * Better message when login required
* **Add option `--print`**: to print any field/template
    * Makes redundant: `--get-description`, `--get-duration`, `--get-filename`, `--get-format`, `--get-id`, `--get-thumbnail`, `--get-title`, `--get-url`
* Field `additional_urls` to download additional videos from metadata using [`--parse-metadata`](https://github.com/yt-dlp/yt-dlp#modifying-metadata)
* Merge youtube-dl: Upto [commit/dfbbe29](https://github.com/ytdl-org/youtube-dl/commit/dfbbe2902fc67f0f93ee47a8077c148055c67a9b)
* Write thumbnail of playlist and add `pl_thumbnail` outtmpl key
* [embedthumbnail] Add `flac` support and refactor `mutagen` code by [pukkandan](https://github.com/pukkandan) and [tripulse](https://github.com/tripulse)
* [audius:artist] Add extractor by [king-millez](https://github.com/king-millez)
* [parlview] Add extractor by [king-millez](https://github.com/king-millez)
* [tenplay] Fix extractor by [king-millez](https://github.com/king-millez)
* [rmcdecouverte] Generalize `_VALID_URL`
* Add compat-option `no-attach-infojson`
* Add field `name` for subtitles
* Ensure `post_extract` and `pre_process` only run once
* Fix `--check-formats` when there is network error
* Standardize `write_debug` and `get_param`
* [options] Alias `--write-comments`, `--no-write-comments`
* [options] Refactor callbacks
* [test:download] Only extract enough videos for `playlist_mincount`
* [extractor] bugfix for when `compat_opts` is not given
* [build] Fix x86 build by [shirt](https://github.com/shirt-dev)
* [cleanup] code formatting, youtube tests and readme

### 2021.05.11
* **Deprecate support for python versions < 3.6**
* **Subtitle extraction from manifests** by [fstirlitz](https://github.com/fstirlitz). See [be6202f](https://github.com/yt-dlp/yt-dlp/commit/be6202f12b97858b9d716e608394b51065d0419f) for details
* **Improve output template:**
    * Allow slicing lists/strings using `field.start:end:step`
    * A field can also be used as offset like `field1+num+field2`
    * A default value can be given using `field|default`
    * Prevent invalid fields from causing errors
* **Merge youtube-dl**: Upto [commit/a726009](https://github.com/ytdl-org/youtube-dl/commit/a7260099873acc6dc7d76cafad2f6b139087afd0)
* **Remove options** `-l`, `-t`, `-A` completely and disable `--auto-number`, `--title`, `--literal`, `--id`
* [Plugins] Prioritize plugins over standard extractors and prevent plugins from overwriting the standard extractor classes
* [downloader] Fix `quiet` and `to_stderr`
* [fragment] Ensure the file is closed on error
* [fragment] Make sure first segment is not skipped
* [aria2c] Fix whitespace being stripped off
* [embedthumbnail] Fix bug where jpeg thumbnails were converted again
* [FormatSort] Fix for when some formats have quality and others don't
* [utils] Add `network_exceptions`
* [utils] Escape URL while sanitizing
* [ukcolumn] Add Extractor
* [whowatch] Add extractor by [nao20010128nao](https://github.com/nao20010128nao)
* [CBS] Improve `_VALID_URL` to support movies
* [crackle] Improve extraction
* [curiositystream] Fix collections
* [francetvinfo] Improve video id extraction
* [generic] Respect the encoding in manifest
* [limelight] Obey `allow_unplayable_formats`
* [mediasite] Generalize URL pattern by [fstirlitz](https://github.com/fstirlitz)
* [mxplayer] Add MxplayerShowIE by [Ashish0804](https://github.com/Ashish0804)
* [nebula] Move to nebula.app by [Lamieur](https://github.com/Lamieur)
* [niconico] Fix HLS formats by [CXwudi](https://github.com/CXwudi), [tsukumijima](https://github.com/tsukumijima), [nao20010128nao](https://github.com/nao20010128nao) and [pukkandan](https://github.com/pukkandan)
* [niconico] Fix title and thumbnail extraction by [CXwudi](https://github.com/CXwudi)
* [plutotv] Extract subtitles from manifests
* [plutotv] Fix format extraction for some urls
* [rmcdecouverte] Improve `_VALID_URL`
* [sonyliv] Fix `title` and `series` extraction by [Ashish0804](https://github.com/Ashish0804)
* [tubi] Raise "no video formats" error when video url is empty
* [youtube:tab] Detect playlists inside community posts
* [youtube] Add `oembed` to reserved names
* [zee5] Fix extraction for some URLs by [Hadi0609](https://github.com/Hadi0609)
* [zee5] Fix py2 compatibility
* Fix `playlist_index` and add `playlist_autonumber`. See [#302](https://github.com/yt-dlp/yt-dlp/issues/302) for details
* Add experimental option `--check-formats` to test the URLs before format selection
* Option `--compat-options` to revert [some of yt-dlp's changes](https://github.com/yt-dlp/yt-dlp#differences-in-default-behavior)
    * Deprecates `--list-formats-as-table`, `--list-formats-old`
* Fix number of digits in `%(playlist_index)s`
* Fix case sensitivity of format selector
* Revert "[core] be able to hand over id and title using url_result"
* Do not strip out whitespaces in `-o` and `-P`
* Fix `preload_download_archive` writing verbose message to `stdout`
* Move option warnings to `YoutubeDL`so that they obey `--no-warnings` and can output colors
* Py2 compatibility for `FileNotFoundError`


### 2021.04.22
* **Improve output template:**
    * Objects can be traversed like `%(field.key1.key2)s`
    * An offset can be added to numeric fields as `%(field+N)s`
    * Deprecates `--autonumber-start`
* **Improve `--sub-langs`:**
    * Treat `--sub-langs` entries as regex
    * `all` can be used to refer to all the subtitles
    * language codes can be prefixed with `-` to exclude it
    * Deprecates `--all-subs`
* Add option `--ignore-no-formats-error` to ignore the "no video format" and similar errors
* Add option `--skip-playlist-after-errors` to skip the rest of a playlist after a given number of errors are encountered
* Merge youtube-dl: Upto [commit/7e8b3f9](https://github.com/ytdl-org/youtube-dl/commit/7e8b3f9439ebefb3a3a4e5da9c0bd2b595976438)
* [downloader] Fix bug in downloader selection
* [BilibiliChannel] Fix pagination by [nao20010128nao](https://github.com/nao20010128nao) and [pukkandan](https://github.com/pukkandan)
* [rai] Add support for http formats by [nixxo](https://github.com/nixxo)
* [TubiTv] Add TubiTvShowIE by [Ashish0804](https://github.com/Ashish0804)
* [twitcasting] Fix extractor
* [viu:ott] Fix extractor and support series by [lkho](https://github.com/lkho) and [pukkandan](https://github.com/pukkandan)
* [youtube:tab] Show unavailable videos in playlists by [coletdjnz](https://github.com/coletdjnz)
* [youtube:tab] Reload with unavailable videos for all playlists
* [youtube] Ignore invalid stretch ratio
* [youtube] Improve channel syncid extraction to support ytcfg by [coletdjnz](https://github.com/coletdjnz)
* [youtube] Standardize API calls for tabs, mixes and search by [coletdjnz](https://github.com/coletdjnz)
* [youtube] Bugfix in `_extract_ytcfg`
* [mildom:user:vod] Download only necessary amount of pages
* [mildom] Remove proxy completely by [fstirlitz](https://github.com/fstirlitz)
* [go] Fix `_VALID_URL`
* [MetadataFromField] Improve regex and add tests
* [Exec] Ensure backward compatibility when the command contains `%`
* [extractor] Fix inconsistent use of `report_warning`
* Ensure `mergeall` selects best format when multistreams are disabled
* Improve the yt-dlp.sh script by [fstirlitz](https://github.com/fstirlitz)
* [lazy_extractor] Do not load plugins
* [ci] Disable fail-fast
* [docs] Clarify which deprecated options still work
* [docs] Fix typos


### 2021.04.11
* Add option `--convert-thumbnails` (only jpg currently supported)
* Format selector `mergeall` to download and merge all formats
* Pass any field to `--exec` using similar syntax to output template
* Choose downloader for each protocol using `--downloader PROTO:NAME`
    * Alias `--downloader` for `--external-downloader`
    * Added `native` as an option for the downloader
* Merge youtube-dl: Upto [commit/4fb25ff](https://github.com/ytdl-org/youtube-dl/commit/4fb25ff5a3be5206bb72e5c4046715b1529fb2c7) (except vimeo)
* [DiscoveryPlusIndia] Add DiscoveryPlusIndiaShowIE by [Ashish0804](https://github.com/Ashish0804)
* [NFHSNetwork] Add extractor by [llacb47](https://github.com/llacb47)
* [nebula] Add extractor (watchnebula.com) by [hheimbuerger](https://github.com/hheimbuerger)
* [nitter] Fix extraction of reply tweets and update instance list by [B0pol](https://github.com/B0pol)
* [nitter] Fix thumbnails by [B0pol](https://github.com/B0pol)
* [youtube] Fix thumbnail URL
* [youtube] Parse API parameters from initial webpage by [coletdjnz](https://github.com/coletdjnz)
* [youtube] Extract comments' approximate timestamp by [coletdjnz](https://github.com/coletdjnz)
* [youtube] Fix alert extraction
* [bilibili] Fix uploader
* [utils] Add `datetime_from_str` and `datetime_add_months` by [coletdjnz](https://github.com/coletdjnz)
* Run some `postprocessors` before actual download
* Improve argument parsing for `-P`, `-o`, `-S`
* Fix some `m3u8` not obeying `--allow-unplayable-formats`
* Fix default of `dynamic_mpd`
* Deprecate `--all-formats`, `--include-ads`, `--hls-prefer-native`, `--hls-prefer-ffmpeg`
* [docs] Improvements

### 2021.04.03
* Merge youtube-dl: Upto [commit/654b4f4](https://github.com/ytdl-org/youtube-dl/commit/654b4f4ff2718f38b3182c1188c5d569c14cc70a)
* Ability to set a specific field in the file's metadata using `--parse-metadata`
* Ability to select n'th best format like `-f bv*.2`
* [DiscoveryPlus] Add discoveryplus.in
* [la7] Add podcasts and podcast playlists by [nixxo](https://github.com/nixxo)
* [mildom] Update extractor with current proxy by [nao20010128nao](https://github.com/nao20010128nao)
* [ard:mediathek] Fix video id extraction
* [generic] Detect Invidious' link element
* [youtube] Show premium state in `availability` by [coletdjnz](https://github.com/coletdjnz)
* [viewsource] Add extractor to handle `view-source:`
* [sponskrub] Run before embedding thumbnail
* [docs] Improve `--parse-metadata` documentation


### 2021.03.24.1
* Revert [commit/8562218](https://github.com/ytdl-org/youtube-dl/commit/8562218350a79d4709da8593bb0c538aa0824acf)

### 2021.03.24
* Merge youtube-dl: Upto 2021.03.25 ([commit/8562218](https://github.com/ytdl-org/youtube-dl/commit/8562218350a79d4709da8593bb0c538aa0824acf))
* Parse metadata from multiple fields using `--parse-metadata`
* Ability to load playlist infojson using `--load-info-json`
* Write current epoch to infojson when using `--no-clean-infojson`
* [youtube_live_chat] fix bug when trying to set cookies
* [niconico] Fix for when logged in by [CXwudi](https://github.com/CXwudi) and [xtkoba](https://github.com/xtkoba)
* [linuxacadamy] Fix login


### 2021.03.21
* Merge youtube-dl: Upto [commit/7e79ba7](https://github.com/ytdl-org/youtube-dl/commit/7e79ba7dd6e6649dd2ce3a74004b2044f2182881)
* Option `--no-clean-infojson` to keep private keys in the infojson
* [aria2c] Support retry/abort unavailable fragments by [damianoamatruda](https://github.com/damianoamatruda)
* [aria2c] Better default arguments
* [movefiles] Fix bugs and make more robust
* [formatSort] Fix `quality` being ignored
* [splitchapters] Fix for older ffmpeg
* [sponskrub] Pass proxy to sponskrub
* Make sure `post_hook` gets the final filename
* Recursively remove any private keys from infojson
* Embed video URL metadata inside `mp4` by [damianoamatruda](https://github.com/damianoamatruda) and [pukkandan](https://github.com/pukkandan)
* Merge `webm` formats into `mkv` if thumbnails are to be embedded by [damianoamatruda](https://github.com/damianoamatruda)
* Use headers and cookies when downloading subtitles by [damianoamatruda](https://github.com/damianoamatruda)
* Parse resolution in info dictionary by [damianoamatruda](https://github.com/damianoamatruda)
* More consistent warning messages by [damianoamatruda](https://github.com/damianoamatruda) and [pukkandan](https://github.com/pukkandan)
* [docs] Add deprecated options and aliases in readme
* [docs] Fix some minor mistakes

* [niconico] Partial fix adapted from [animelover1984/youtube-dl@b5eff52](https://github.com/animelover1984/youtube-dl/commit/b5eff52dd9ed5565672ea1694b38c9296db3fade) (login and smile formats still don't work)
* [niconico] Add user extractor by [animelover1984](https://github.com/animelover1984)
* [bilibili] Add anthology support by [animelover1984](https://github.com/animelover1984)
* [amcnetworks] Fix extractor by [2ShedsJackson](https://github.com/2ShedsJackson)
* [stitcher] Merge from youtube-dl by [nixxo](https://github.com/nixxo)
* [rcs] Improved extraction by [nixxo](https://github.com/nixxo)
* [linuxacadamy] Improve regex
* [youtube] Show if video is `private`, `unlisted` etc in info (`availability`) by [coletdjnz](https://github.com/coletdjnz) and [pukkandan](https://github.com/pukkandan)
* [youtube] bugfix for channel playlist extraction
* [nbc] Improve metadata extraction by [2ShedsJackson](https://github.com/2ShedsJackson)


### 2021.03.15
* **Split video by chapters**: using option `--split-chapters`
    * The output file of the split files can be set with `-o`/`-P` using the prefix `chapter:`
    * Additional keys `section_title`, `section_number`, `section_start`, `section_end` are available in the output template
* **Parallel fragment downloads** by [shirt](https://github.com/shirt-dev)
    * Use option `--concurrent-fragments` (`-N`) to set the number of threads (default 1)
* Merge youtube-dl: Upto [commit/3be0980](https://github.com/ytdl-org/youtube-dl/commit/3be098010f667b14075e3dfad1e74e5e2becc8ea)
* [zee5] Add Show Extractor by [Ashish0804](https://github.com/Ashish0804) and [pukkandan](https://github.com/pukkandan)
* [rai] fix drm check [nixxo](https://github.com/nixxo)
* [wimtv] Add extractor by [nixxo](https://github.com/nixxo)
* [mtv] Add mtv.it and extract series metadata by [nixxo](https://github.com/nixxo)
* [pluto.tv] Add extractor by [kevinoconnor7](https://github.com/kevinoconnor7)
* [youtube] Rewrite comment extraction by [coletdjnz](https://github.com/coletdjnz)
* [embedthumbnail] Set mtime correctly
* Refactor some postprocessor/downloader code by [pukkandan](https://github.com/pukkandan) and [shirt](https://github.com/shirt-dev)


### 2021.03.07
* [youtube] Fix history, mixes, community pages and trending by [pukkandan](https://github.com/pukkandan) and [coletdjnz](https://github.com/coletdjnz)
* [youtube] Fix private feeds/playlists on multi-channel accounts by [coletdjnz](https://github.com/coletdjnz)
* [youtube] Extract alerts from continuation by [coletdjnz](https://github.com/coletdjnz)
* [cbs] Add support for ParamountPlus by [shirt](https://github.com/shirt-dev)
* [mxplayer] Rewrite extractor with show support by [pukkandan](https://github.com/pukkandan) and [Ashish0804](https://github.com/Ashish0804)
* [gedi] Improvements from youtube-dl by [nixxo](https://github.com/nixxo)
* [vimeo] Fix videos with password by [teesid](https://github.com/teesid)
* [lbry] Support `lbry://` url by [nixxo](https://github.com/nixxo)
* [bilibili] Change `Accept` header by [pukkandan](https://github.com/pukkandan) and [animelover1984](https://github.com/animelover1984)
* [trovo] Pass origin header
* [rai] Check for DRM by [nixxo](https://github.com/nixxo)
* [downloader] Fix bug for `ffmpeg`/`httpie`
* [update] Fix updater removing the executable bit on some UNIX distros
* [update] Fix current build hash for UNIX
* [docs] Include wget/curl/aria2c install instructions for Unix by [Ashish0804](https://github.com/Ashish0804)
* Fix some videos downloading with `m3u8` extension
* Remove "fixup is ignored" warning when fixup wasn't passed by user


### 2021.03.03.2
* [build] Fix bug

### 2021.03.03
* [youtube] Use new browse API for continuation page extraction by [coletdjnz](https://github.com/coletdjnz) and [pukkandan](https://github.com/pukkandan)
* Fix HLS playlist downloading by [shirt](https://github.com/shirt-dev)
* Merge youtube-dl: Upto [2021.03.03](https://github.com/ytdl-org/youtube-dl/releases/tag/2021.03.03)
* [mtv] Fix extractor
* [nick] Fix extractor by [DennyDai](https://github.com/DennyDai)
* [mxplayer] Add new extractor by [codeasashu](https://github.com/codeasashu)
* [youtube] Throw error when `--extractor-retries` are exhausted
* Reduce default of `--extractor-retries` to 3
* Fix packaging bugs by [hseg](https://github.com/hseg)


### 2021.03.01
* Allow specifying path in `--external-downloader`
* Add option `--sleep-requests` to sleep b/w requests
* Add option `--extractor-retries` to retry on known extractor errors
* Extract comments only when needed
* `--get-comments` doesn't imply `--write-info-json` if `-J`, `-j` or `--print-json` are used
* Fix `get_executable_path` by [shirt](https://github.com/shirt-dev)
* [youtube] Retry on more known errors than just HTTP-5xx
* [youtube] Fix inconsistent `webpage_url`
* [tennistv] Fix format sorting
* [bilibiliaudio] Recognize the file as audio-only
* [hrfensehen] Fix wrong import
* [viki] Fix viki play pass authentication by [RobinD42](https://github.com/RobinD42)
* [readthedocs] Improvements by [shirt](https://github.com/shirt-dev)
* [hls] Fix bug with m3u8 format extraction
* [hls] Enable `--hls-use-mpegts` by default when downloading live-streams
* [embedthumbnail] Fix bug with deleting original thumbnail
* [build] Fix completion paths, zsh pip completion install by [hseg](https://github.com/hseg)
* [ci] Disable download tests unless specifically invoked
* Cleanup some code and fix typos


### 2021.02.24
* Moved project to an organization [yt-dlp](https://github.com/yt-dlp)
* **Completely changed project name to yt-dlp** by [Pccode66](https://github.com/Pccode66) and [pukkandan](https://github.com/pukkandan)
    * Also, `youtube-dlc` config files are no longer loaded
* Merge youtube-dl: Upto [commit/4460329](https://github.com/ytdl-org/youtube-dl/commit/44603290e5002153f3ebad6230cc73aef42cc2cd) (except tmz, gedi)
* [Readthedocs](https://yt-dlp.readthedocs.io) support by [shirt](https://github.com/shirt-dev)
* [youtube] Show if video was a live stream in info (`was_live`)
* [Zee5] Add new extractor by [Ashish0804](https://github.com/Ashish0804) and [pukkandan](https://github.com/pukkandan)
* [jwplatform] Add support for `hyland.com`
* [tennistv] Fix extractor
* [hls] Support media initialization by [shirt](https://github.com/shirt-dev)
* [hls] Added options `--hls-split-discontinuity` to better support media discontinuity by [shirt](https://github.com/shirt-dev)
* [ffmpeg] Allow passing custom arguments before -i using `--ppa "ffmpeg_i1:ARGS"` syntax
* Fix `--windows-filenames` removing `/` from UNIX paths
* [hls] Show warning if pycryptodome is not found
* [docs] Improvements
    * Fix documentation of `Extractor Options`
    * Document `all` in format selection
    * Document `playable_in_embed` in output templates


### 2021.02.19
* Merge youtube-dl: Upto [commit/cf2dbec](https://github.com/ytdl-org/youtube-dl/commit/cf2dbec6301177a1fddf72862de05fa912d9869d) (except kakao)
* [viki] Fix extractor
* [niconico] Extract `channel` and `channel_id` by [kurumigi](https://github.com/kurumigi)
* [youtube] Multiple page support for hashtag URLs
* [youtube] Add more invidious instances
* [youtube] Fix comment extraction when comment text is empty
* Option `--windows-filenames` to force use of windows compatible filenames
* [ExtractAudio] Bugfix
* Don't raise `parser.error` when exiting for update
* [MoveFiles] Fix for when merger can't run
* Changed `--trim-file-name` to `--trim-filenames` to be similar to related options
* Format Sort improvements:
    * Prefer `vp9.2` more than other `vp9` codecs
    * Remove forced priority of `quality`
    * Remove unnecessary `field_preference` and misuse of `preference` from extractors
* Build improvements:
    * Fix hash output by [shirt](https://github.com/shirt-dev)
    * Lock python package versions for x86 and use `wheels` by [shirt](https://github.com/shirt-dev)
    * Exclude `vcruntime140.dll` from UPX by [jbruchon](https://github.com/jbruchon)
    * Set version number based on UTC time, not local time
    * Publish on PyPi only if token is set
* [docs] Better document `--prefer-free-formats` and add `--no-prefer-free-format`


### 2021.02.15
* Merge youtube-dl: Upto [2021.02.10](https://github.com/ytdl-org/youtube-dl/releases/tag/2021.02.10) (except archive.org)
* [niconico] Improved extraction and support encrypted/SMILE movies by [kurumigi](https://github.com/kurumigi), [tsukumijima](https://github.com/tsukumijima), [bbepis](https://github.com/bbepis), [pukkandan](https://github.com/pukkandan)
* Fix HLS AES-128 with multiple keys in external downloaders by [shirt](https://github.com/shirt-dev)
* [youtube_live_chat] Fix by using POST API by [siikamiika](https://github.com/siikamiika)
* [rumble] Add support for video page
* Option `--allow-unplayable-formats` to allow downloading unplayable video formats
* [ExtractAudio] Don't re-encode when file is already in a common audio format
* [youtube] Fix search continuations
* [youtube] Fix for new accounts
* Improve build/updater: by [pukkandan](https://github.com/pukkandan) and [shirt](https://github.com/shirt-dev)
    * Fix SHA256 calculation in build and implement hash checking for updater
    * Exit immediately in windows once the update process starts
    * Fix updater for `x86.exe`
    * Updater looks for both `yt-dlp` and `youtube-dlc` in releases for future-proofing
    * Change optional dependency to `pycryptodome`
* Fix issue with unicode filenames in aria2c by [shirt](https://github.com/shirt-dev)
* Fix `allow_playlist_files` not being correctly passed through
* Fix for empty HTTP head requests by [shirt](https://github.com/shirt-dev)
* Fix `get_executable_path` in UNIX
* [sponskrub] Print ffmpeg output and errors to terminal
* `__real_download` should be false when ffmpeg unavailable and no download
* Show `exe`/`zip`/`source` and 32/64bit in verbose message


### 2021.02.09
* **aria2c support for DASH/HLS**: by [shirt](https://github.com/shirt-dev)
* **Implement Updater** (`-U`) by [shirt](https://github.com/shirt-dev)
* [youtube] Fix comment extraction
* [youtube_live_chat] Improve extraction
* [youtube] Fix for channel URLs sometimes not downloading all pages
* [aria2c] Changed default arguments to `--console-log-level=warn --summary-interval=0 --file-allocation=none -x16 -j16 -s16`
* Add fallback for thumbnails
* [embedthumbnail] Keep original thumbnail after conversion if write_thumbnail given
* [embedsubtitle] Keep original subtitle after conversion if write_subtitles given
* [pyinst.py] Move back to root dir
* [youtube] Simplified renderer parsing and bugfixes
* [movefiles] Fix compatibility with python2
* [remuxvideo] Fix validation of conditional remux
* [sponskrub] Don't raise error when the video does not exist
* [docs] Crypto is an optional dependency


### 2021.02.04
* Merge youtube-dl: Upto [2021.02.04.1](https://github.com/ytdl-org/youtube-dl/releases/tag/2021.02.04.1)
* **Date/time formatting in output template:**
    * You can use [`strftime`](https://docs.python.org/3/library/datetime.html#strftime-and-strptime-format-codes) to format date/time fields. Example: `%(upload_date>%Y-%m-%d)s`
* **Multiple output templates:**
    * Separate output templates can be given for the different metadata files by using `-o TYPE:TEMPLATE`
    * The allowed types are: `subtitle|thumbnail|description|annotation|infojson|pl_description|pl_infojson`
* [youtube] More metadata extraction for channel/playlist URLs (channel, uploader, thumbnail, tags)
* New option `--no-write-playlist-metafiles` to prevent writing playlist metadata files
* [audius] Fix extractor
* [youtube_live_chat] Fix `parse_yt_initial_data` and add `fragment_retries`
* [postprocessor] Raise errors correctly
* [metadatafromtitle] Fix bug when extracting data from numeric fields
* Fix issue with overwriting files
* Fix "Default format spec" appearing in quiet mode
* [FormatSort] Allow user to prefer av01 over vp9 (The default is still vp9)
* [FormatSort] fix bug where `quality` had more priority than `hasvid`
* [pyinst] Automatically detect python architecture and working directory
* Strip out internal fields such as `_filename` from infojson


### 2021.01.29
* **Features from [animelover1984/youtube-dl](https://github.com/animelover1984/youtube-dl)**: by [animelover1984](https://github.com/animelover1984) and [bbepis](https://github.com/bbepis)
    * Add `--get-comments`
    * [youtube] Extract comments
    * [billibilli] Added BiliBiliSearchIE, BilibiliChannelIE
    * [billibilli] Extract comments
    * [billibilli] Better video extraction
    * Write playlist data to infojson
    * [FFmpegMetadata] Embed infojson inside the video
    * [EmbedThumbnail] Try embedding in mp4 using ffprobe and `-disposition`
    * [EmbedThumbnail] Treat mka like mkv and mov like mp4
    * [EmbedThumbnail] Embed in ogg/opus
    * [VideoRemuxer] Conditionally remux video
    * [VideoRemuxer] Add `-movflags +faststart` when remuxing to mp4
    * [ffmpeg] Print entire stderr in verbose when there is error
    * [EmbedSubtitle] Warn when embedding ass in mp4
    * [anvato] Use NFLTokenGenerator if possible
* **Parse additional metadata**: New option `--parse-metadata` to extract additional metadata from existing fields
    * The extracted fields can be used in `--output`
    * Deprecated `--metadata-from-title`
* [Audius] Add extractor
* [youtube] Extract playlist description and write it to `.description` file
* Detect existing files even when using `recode`/`remux` (`extract-audio` is partially fixed)
* Fix wrong user config from v2021.01.24
* [youtube] Report error message from youtube as error instead of warning
* [FormatSort] Fix some fields not sorting from v2021.01.24
* [postprocessor] Deprecate `avconv`/`avprobe`.  All current functionality is left untouched. But don't expect any new features to work with avconv
* [postprocessor] fix `write_debug` to not throw error when there is no `_downloader`
* [movefiles] Don't give "cant find" warning when move is unnecessary
* Refactor `update-version`, `pyinst.py` and related files
* [ffmpeg] Document more formats that are supported for remux/recode


### 2021.01.24
* Merge youtube-dl: Upto [2021.01.24](https://github.com/ytdl-org/youtube-dl/releases/tag/2021.01.16)
* Plugin support ([documentation](https://github.com/yt-dlp/yt-dlp#plugins))
* **Multiple paths**: New option `-P`/`--paths` to give different paths for different types of files
    * The syntax is `-P "type:path" -P "type:path"`
    * Valid types are: home, temp, description, annotation, subtitle, infojson, thumbnail
    * Additionally, configuration file is taken from home directory or current directory
* Allow passing different arguments to different external downloaders
* [mildom] Add extractor by [nao20010128nao](https://github.com/nao20010128nao)
* Warn when using old style `--external-downloader-args` and `--post-processor-args`
* Fix `--no-overwrite` when using `--write-link`
* [sponskrub] Output `unrecognized argument` error message correctly
* [cbs] Make failure to extract title non-fatal
* Fix typecasting when pre-checking archive
* Fix issue with setting title on UNIX
* Deprecate redundant aliases in `formatSort`. The aliases remain functional for backward compatibility, but will be left undocumented
* [tests] Fix test_post_hooks
* [tests] Split core and download tests


### 2021.01.20
* [TrovoLive] Add extractor (only VODs)
* [pokemon] Add `/#/player` URLs
* Improved parsing of multiple postprocessor-args, add `--ppa` as alias
* [EmbedThumbnail] Simplify embedding in mkv
* [sponskrub] Encode filenames correctly, better debug output and error message
* [readme] Cleanup options


### 2021.01.16
* Merge youtube-dl: Upto [2021.01.16](https://github.com/ytdl-org/youtube-dl/releases/tag/2021.01.16)
* **Configuration files:**
    * Portable configuration file: `./yt-dlp.conf`
    * Allow the configuration files to be named `yt-dlp` instead of `youtube-dlc`. See [this](https://github.com/yt-dlp/yt-dlp#configuration) for details
* Add PyPI release


### 2021.01.14
* Added option `--break-on-reject`
* [roosterteeth.com] Fix for bonus episodes by [Zocker1999NET](https://github.com/Zocker1999NET)
* [tiktok] Fix for when share_info is empty
* [EmbedThumbnail] Fix bug due to incorrect function name
* [docs] Changed sponskrub links to point to [yt-dlp/SponSkrub](https://github.com/yt-dlp/SponSkrub) since I am now providing both linux and windows releases
* [docs] Change all links to correctly point to new fork URL
* [docs] Fixes typos


### 2021.01.12
* [roosterteeth.com] Add subtitle support by [samiksome](https://github.com/samiksome)
* Added `--force-overwrites`, `--no-force-overwrites` by [alxnull](https://github.com/alxnull)
* Changed fork name to `yt-dlp`
* Fix typos by [FelixFrog](https://github.com/FelixFrog)
* [ci] Option to skip
* [changelog] Added unreleased changes in blackjack4494/yt-dlc


### 2021.01.10
* [archive.org] Fix extractor and add support for audio and playlists by [wporr](https://github.com/wporr)
* [Animelab] Added by [mariuszskon](https://github.com/mariuszskon)
* [youtube:search] Fix view_count by [ohnonot](https://github.com/ohnonot)
* [youtube] Show if video is embeddable in info (`playable_in_embed`)
* Update version badge automatically in README
* Enable `test_youtube_search_matching`
* Create `to_screen` and similar functions in postprocessor/common


### 2021.01.09
* [youtube] Fix bug in automatic caption extraction
* Add `post_hooks` to YoutubeDL by [alexmerkel](https://github.com/alexmerkel)
* Batch file enumeration improvements by [glenn-slayden](https://github.com/glenn-slayden)
* Stop immediately when reaching `--max-downloads` by [glenn-slayden](https://github.com/glenn-slayden)
* Fix incorrect ANSI sequence for restoring console-window title by [glenn-slayden](https://github.com/glenn-slayden)
* Kill child processes when yt-dlc is killed by [Unrud](https://github.com/Unrud)


### 2021.01.08
* Merge youtube-dl: Upto [2021.01.08](https://github.com/ytdl-org/youtube-dl/releases/tag/2021.01.08) except stitcher ([1](https://github.com/ytdl-org/youtube-dl/commit/bb38a1215718cdf36d73ff0a7830a64cd9fa37cc), [2](https://github.com/ytdl-org/youtube-dl/commit/a563c97c5cddf55f8989ed7ea8314ef78e30107f))
* Moved changelog to separate file


### 2021.01.07-1
* [Akamai] fix by [nixxo](https://github.com/nixxo)
* [Tiktok] merge youtube-dl tiktok extractor by [GreyAlien502](https://github.com/GreyAlien502)
* [vlive] add support for playlists by [kyuyeunk](https://github.com/kyuyeunk)
* [youtube_live_chat] make sure playerOffsetMs is positive by [siikamiika](https://github.com/siikamiika)
* Ignore extra data streams in ffmpeg by [jbruchon](https://github.com/jbruchon)
* Allow passing different arguments to different postprocessors using `--postprocessor-args`
* Deprecated `--sponskrub-args`. The same can now be done using `--postprocessor-args "sponskrub:<args>"`
* [CI] Split tests into core-test and full-test


### 2021.01.07
* Removed priority of `av01` codec in `-S` since most devices don't support it yet
* Added `duration_string` to be used in `--output`
* Created First Release


### 2021.01.05-1
* **Changed defaults:**
    * Enabled `--ignore`
    * Disabled `--video-multistreams` and `--audio-multistreams`
    * Changed default format selection to `bv*+ba/b` when `--audio-multistreams` is disabled
    * Changed default format sort order to `res,fps,codec,size,br,asr,proto,ext,has_audio,source,format_id`
    * Changed `webm` to be more preferable than `flv` in format sorting
    * Changed default output template to `%(title)s [%(id)s].%(ext)s`
    * Enabled `--list-formats-as-table`


### 2021.01.05
* **Format Sort:** Added `--format-sort` (`-S`), `--format-sort-force` (`--S-force`) - See [Sorting Formats](README.md#sorting-formats) for details
* **Format Selection:** See [Format Selection](README.md#format-selection) for details
    * New format selectors: `best*`, `worst*`, `bestvideo*`, `bestaudio*`, `worstvideo*`, `worstaudio*`
    * Changed video format sorting to show video only files and video+audio files together
    * Added `--video-multistreams`, `--no-video-multistreams`, `--audio-multistreams`, `--no-audio-multistreams`
    * Added `b`,`w`,`v`,`a` as alias for `best`, `worst`, `video` and `audio` respectively
* Shortcut Options: Added `--write-link`, `--write-url-link`, `--write-webloc-link`, `--write-desktop-link` by [h-h-h-h](https://github.com/h-h-h-h) - See [Internet Shortcut Options](README.md#internet-shortcut-options) for details
* **Sponskrub integration:** Added `--sponskrub`, `--sponskrub-cut`, `--sponskrub-force`, `--sponskrub-location`, `--sponskrub-args` - See [SponSkrub Options](README.md#sponskrub-sponsorblock-options) for details
* Added `--force-download-archive` (`--force-write-archive`) by [h-h-h-h](https://github.com/h-h-h-h)
* Added `--list-formats-as-table`,  `--list-formats-old`
* **Negative Options:** Makes it possible to negate most boolean options by adding a `no-` to the switch. Usefull when you want to reverse an option that is defined in a config file
    * Added `--no-ignore-dynamic-mpd`, `--no-allow-dynamic-mpd`, `--allow-dynamic-mpd`, `--youtube-include-hls-manifest`, `--no-youtube-include-hls-manifest`, `--no-youtube-skip-hls-manifest`, `--no-download`, `--no-download-archive`, `--resize-buffer`, `--part`, `--mtime`, `--no-keep-fragments`, `--no-cookies`, `--no-write-annotations`, `--no-write-info-json`, `--no-write-description`, `--no-write-thumbnail`, `--youtube-include-dash-manifest`, `--post-overwrites`, `--no-keep-video`, `--no-embed-subs`, `--no-embed-thumbnail`, `--no-add-metadata`, `--no-include-ads`, `--no-write-sub`, `--no-write-auto-sub`, `--no-playlist-reverse`, `--no-restrict-filenames`, `--youtube-include-dash-manifest`, `--no-format-sort-force`, `--flat-videos`, `--no-list-formats-as-table`, `--no-sponskrub`, `--no-sponskrub-cut`, `--no-sponskrub-force`
    * Renamed: `--write-subs`, `--no-write-subs`, `--no-write-auto-subs`, `--write-auto-subs`. Note that these can still be used without the ending "s"
* Relaxed validation for format filters so that any arbitrary field can be used
* Fix for embedding thumbnail in mp3 by [pauldubois98](https://github.com/pauldubois98) ([ytdl-org/youtube-dl#21569](https://github.com/ytdl-org/youtube-dl/pull/21569))
* Make Twitch Video ID output from Playlist and VOD extractor same. This is only a temporary fix
* Merge youtube-dl: Upto [2021.01.03](https://github.com/ytdl-org/youtube-dl/commit/8e953dcbb10a1a42f4e12e4e132657cb0100a1f8) - See [blackjack4494/yt-dlc#280](https://github.com/blackjack4494/yt-dlc/pull/280) for details
    * Extractors [tiktok](https://github.com/ytdl-org/youtube-dl/commit/fb626c05867deab04425bad0c0b16b55473841a2) and [hotstar](https://github.com/ytdl-org/youtube-dl/commit/bb38a1215718cdf36d73ff0a7830a64cd9fa37cc) have not been merged
* Cleaned up the fork for public use


**Note**: All uncredited changes above this point are authored by [pukkandan](https://github.com/pukkandan)

### Unreleased changes in [blackjack4494/yt-dlc](https://github.com/blackjack4494/yt-dlc)
* Updated to youtube-dl release 2020.11.26 by [pukkandan](https://github.com/pukkandan)
* Youtube improvements by [pukkandan](https://github.com/pukkandan)
    * Implemented all Youtube Feeds (ytfav, ytwatchlater, ytsubs, ythistory, ytrec) and SearchURL
    * Fix some improper Youtube URLs
    * Redirect channel home to /video
    * Print youtube's warning message
    * Handle Multiple pages for feeds better
* [youtube] Fix ytsearch not returning results sometimes due to promoted content by [coletdjnz](https://github.com/coletdjnz)
* [youtube] Temporary fix for automatic captions - disable json3 by [blackjack4494](https://github.com/blackjack4494)
* Add --break-on-existing by [gergesh](https://github.com/gergesh)
* Pre-check video IDs in the archive before downloading by [pukkandan](https://github.com/pukkandan)
* [bitwave.tv] New extractor by [lorpus](https://github.com/lorpus)
* [Gedi] Add extractor by [nixxo](https://github.com/nixxo)
* [Rcs] Add new extractor by [nixxo](https://github.com/nixxo)
* [skyit] New skyitalia extractor by [nixxo](https://github.com/nixxo)
* [france.tv] Fix thumbnail URL by [renalid](https://github.com/renalid)
* [ina] support mobile links by [B0pol](https://github.com/B0pol)
* [instagram] Fix thumbnail extractor by [nao20010128nao](https://github.com/nao20010128nao)
* [SouthparkDe] Support for English URLs by [xypwn](https://github.com/xypwn)
* [spreaker] fix SpreakerShowIE test URL by [pukkandan](https://github.com/pukkandan)
* [Vlive] Fix playlist handling when downloading a channel by [kyuyeunk](https://github.com/kyuyeunk)
* [tmz] Fix extractor by [diegorodriguezv](https://github.com/diegorodriguezv)
* [ITV] BTCC URL update by [WolfganP](https://github.com/WolfganP)
* [generic] Detect embedded bitchute videos by [pukkandan](https://github.com/pukkandan)
* [generic] Extract embedded youtube and twitter videos by [diegorodriguezv](https://github.com/diegorodriguezv)
* [ffmpeg] Ensure all streams are copied by [pukkandan](https://github.com/pukkandan)
* [embedthumbnail] Fix for os.rename error by [pukkandan](https://github.com/pukkandan)
* make_win.bat: don't use UPX to pack vcruntime140.dll by [jbruchon](https://github.com/jbruchon)


### Changelog of [blackjack4494/yt-dlc](https://github.com/blackjack4494/yt-dlc) till release 2020.11.11-3

**Note**: This was constructed from the merge commit messages and may not be entirely accurate

* [bandcamp] fix failing test. remove subclass hack by [insaneracist](https://github.com/insaneracist)
* [bandcamp] restore album downloads by [insaneracist](https://github.com/insaneracist)
* [francetv] fix extractor by [Surkal](https://github.com/Surkal)
* [gdcvault] fix extractor by [blackjack4494](https://github.com/blackjack4494)
* [hotstar] Move to API v1 by [theincognito-inc](https://github.com/theincognito-inc)
* [hrfernsehen] add extractor by [blocktrron](https://github.com/blocktrron)
* [kakao] new apis by [blackjack4494](https://github.com/blackjack4494)
* [la7] fix missing protocol by [nixxo](https://github.com/nixxo)
* [mailru] removed escaped braces, use urljoin, added tests by [nixxo](https://github.com/nixxo)
* [MTV/Nick] universal mgid extractor + fix nick.de feed by [blackjack4494](https://github.com/blackjack4494)
* [mtv] Fix a missing match_id by [nixxo](https://github.com/nixxo)
* [Mtv] updated extractor logic & more by [blackjack4494](https://github.com/blackjack4494)
* [ndr] support Daserste ndr by [blackjack4494](https://github.com/blackjack4494)
* [Netzkino] Only use video id to find metadata by [TobiX](https://github.com/TobiX)
* [newgrounds] fix: video download by [insaneracist](https://github.com/insaneracist)
* [nitter] Add new extractor by [B0pol](https://github.com/B0pol)
* [soundcloud] Resolve audio/x-wav by [tfvlrue](https://github.com/tfvlrue)
* [soundcloud] sets pattern and tests by [blackjack4494](https://github.com/blackjack4494)
* [SouthparkDE/MTV] another mgid extraction (mtv_base) feed url updated by [blackjack4494](https://github.com/blackjack4494)
* [StoryFire] Add new extractor by [sgstair](https://github.com/sgstair)
* [twitch] by [geauxlo](https://github.com/geauxlo)
* [videa] Adapt to updates by [adrianheine](https://github.com/adrianheine)
* [Viki] subtitles, formats by [blackjack4494](https://github.com/blackjack4494)
* [vlive] fix extractor for revamped website by [exwm](https://github.com/exwm)
* [xtube] fix extractor by [insaneracist](https://github.com/insaneracist)
* [youtube] Convert subs when download is skipped by [blackjack4494](https://github.com/blackjack4494)
* [youtube] Fix age gate detection by [random-nick](https://github.com/random-nick)
* [youtube] fix yt-only playback when age restricted/gated - requires cookies by [blackjack4494](https://github.com/blackjack4494)
* [youtube] fix: extract artist metadata from ytInitialData by [insaneracist](https://github.com/insaneracist)
* [youtube] fix: extract mix playlist ids from ytInitialData by [insaneracist](https://github.com/insaneracist)
* [youtube] fix: mix playlist title by [insaneracist](https://github.com/insaneracist)
* [youtube] fix: Youtube Music playlists by [insaneracist](https://github.com/insaneracist)
* [Youtube] Fixed problem with new youtube player by [peet1993](https://github.com/peet1993)
* [zoom] Fix url parsing for url's containing /share/ and dots by [Romern](https://github.com/Romern)
* [zoom] new extractor by [insaneracist](https://github.com/insaneracist)
* abc by [adrianheine](https://github.com/adrianheine)
* Added Comcast_SSO fix by [merval](https://github.com/merval)
* Added DRM logic to brightcove by [merval](https://github.com/merval)
* Added regex for ABC.com site. by [kucksdorfs](https://github.com/kucksdorfs)
* alura by [hugohaa](https://github.com/hugohaa)
* Arbitrary merges by [fstirlitz](https://github.com/fstirlitz)
* ard.py_add_playlist_support by [martin54](https://github.com/martin54)
* Bugfix/youtube/chapters fix extractor by [gschizas](https://github.com/gschizas)
* bugfix_youtube_like_extraction by [RedpointsBots](https://github.com/RedpointsBots)
* Create build workflow by [blackjack4494](https://github.com/blackjack4494)
* deezer by [LucBerge](https://github.com/LucBerge)
* Detect embedded bitchute videos by [pukkandan](https://github.com/pukkandan)
* Don't install tests by [l29ah](https://github.com/l29ah)
* Don't try to embed/convert json subtitles generated by [youtube](https://github.com/youtube) livechat by [pukkandan](https://github.com/pukkandan)
* Doodstream by [sxvghd](https://github.com/sxvghd)
* duboku by [lkho](https://github.com/lkho)
* elonet by [tpikonen](https://github.com/tpikonen)
* ext/remuxe-video by [Zocker1999NET](https://github.com/Zocker1999NET)
* fall-back to the old way to fetch subtitles, if needed by [RobinD42](https://github.com/RobinD42)
* feature_subscriber_count by [RedpointsBots](https://github.com/RedpointsBots)
* Fix external downloader when there is no http_header by [pukkandan](https://github.com/pukkandan)
* Fix issue triggered by [tubeup](https://github.com/tubeup) by [nsapa](https://github.com/nsapa)
* Fix YoutubePlaylistsIE by [ZenulAbidin](https://github.com/ZenulAbidin)
* fix-mitele' by [DjMoren](https://github.com/DjMoren)
* fix/google-drive-cookie-issue by [legraphista](https://github.com/legraphista)
* fix_tiktok by [mervel-mervel](https://github.com/mervel-mervel)
* Fixed problem with JS player URL by [peet1993](https://github.com/peet1993)
* fixYTSearch by [xarantolus](https://github.com/xarantolus)
* FliegendeWurst-3sat-zdf-merger-bugfix-feature
* gilou-bandcamp_update
* implement ThisVid extractor by [rigstot](https://github.com/rigstot)
* JensTimmerman-patch-1 by [JensTimmerman](https://github.com/JensTimmerman)
* Keep download archive in memory for better performance by [jbruchon](https://github.com/jbruchon)
* la7-fix by [iamleot](https://github.com/iamleot)
* magenta by [adrianheine](https://github.com/adrianheine)
* Merge 26564 from [adrianheine](https://github.com/adrianheine)
* Merge code from [ddland](https://github.com/ddland)
* Merge code from [nixxo](https://github.com/nixxo)
* Merge code from [ssaqua](https://github.com/ssaqua)
* Merge code from [zubearc](https://github.com/zubearc)
* mkvthumbnail by [MrDoritos](https://github.com/MrDoritos)
* myvideo_ge by [fonkap](https://github.com/fonkap)
* naver by [SeonjaeHyeon](https://github.com/SeonjaeHyeon)
* ondemandkorea by [julien-hadleyjack](https://github.com/julien-hadleyjack)
* rai-update by [iamleot](https://github.com/iamleot)
* RFC: youtube: Polymer UI and JSON endpoints for playlists by [wlritchi](https://github.com/wlritchi)
* rutv by [adrianheine](https://github.com/adrianheine)
* Sc extractor web auth by [blackjack4494](https://github.com/blackjack4494)
* Switch from binary search tree to Python sets by [jbruchon](https://github.com/jbruchon)
* tiktok by [skyme5](https://github.com/skyme5)
* tvnow by [TinyToweringTree](https://github.com/TinyToweringTree)
* twitch-fix by [lel-amri](https://github.com/lel-amri)
* Twitter shortener by [blackjack4494](https://github.com/blackjack4494)
* Update README.md by [JensTimmerman](https://github.com/JensTimmerman)
* Update to reflect website changes. by [amigatomte](https://github.com/amigatomte)
* use webarchive to fix a dead link in README by [B0pol](https://github.com/B0pol)
* Viki the second by [blackjack4494](https://github.com/blackjack4494)
* wdr-subtitles by [mrtnmtth](https://github.com/mrtnmtth)
* Webpfix by [alexmerkel](https://github.com/alexmerkel)
* Youtube live chat by [siikamiika](https://github.com/siikamiika)
