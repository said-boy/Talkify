
![Talkify](https://github.com/said-boy/Talkify/assets/97724397/6ab747a0-9b90-496a-b996-0edfcafe5a20)


# 
**Talkify** adalah sebuah aplikasi yang didesain **khusus untuk warga negara Indonesia** terkhusus jika anda mungkin malu untuk berbicara kepada orang secara langsung. Aplikasi ini memungkinkan pengguna untuk berkomunikasi dengan kecerdasan buatan melalui suara dalam bahasa Inggris. 

Melalui Talkify, Anda dapat memantau tingkat level dari bahasa inggris anda dan Anda juga dapat meningkatkan kemampuan tata bahasa dan pengucapan Bahasa Inggris Anda secara efektif. Program ini dirancang untuk membantu Anda meraih tingkat kefasihan yang lebih tinggi dalam berbahasa Inggris dengan lebih percaya diri.




## Persyaratan

- Linux
- sox (aplikasi perekam suara di linux)
- Golang v.1.22.0
- Api key wit
- Api key gemini


## Melengkapi persyaratan

- Download **sox** untuk Debian & Ubuntu

```bash
  $ sudo apt install sox
```

- Cara install **Golang** klik [disini](https://go.dev/doc/install)
- Mendapatkan api key **Wit** ada [disini](https://wit.ai/) , silahkan daftar dan dapatkan api key nya.
- Mendapatkan api key **Gemini** ada [disini](https://ai.google.dev/?gad_source=1&gclid=Cj0KCQiA5-uuBhDzARIsAAa21T-QwHqAsQxRms4uHy1vgYSztCRL5ihdxqPVyL4NBFGb6zAZZCmhBqkaAj34EALw_wcB) , silahkan daftar dan dapatkan api key nya.
    
## Menjalankan Aplikasi

Download project

```bash
  git clone https://github.com/said-boy/Talkify.git
```

Masuk ke folder Talkify

```bash
  cd Talkify
```

Install dependensi

```bash
  go mod tidy
```

Setting variable environment 

- rename file config.yaml.conf yang ada di folder env/ menjadi config.yaml
- masukkan `API_KEY` kedalam file config.yaml tersebut.

Menjalankan aplikasi

```bash
  go run cmd/main.go
```


## Tampilan aplikasi

![Talkify_SS](https://github.com/said-boy/Talkify/assets/97724397/f4b3bac0-1af1-410d-b6b0-c360d077b273)


## Contribusi

Jika anda ingin berkontribusi dalam pengembangan projek ini.

 silahkan bisa melakukannya melalui pull request


## Lisensi


[MIT License](https://choosealicense.com/licenses/mit/)

tes
