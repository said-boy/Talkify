import re
import speech_recognition as sr
from gtts import gTTS
import os
import requests
import json

def getAnswer(text) -> str: 
    API_KEY = "YOUR_GEMINI_API_KEY"

    url = "https://generativelanguage.googleapis.com/v1beta/models/gemini-pro:generateContent?key=" + API_KEY

    data = {
        "contents": [{
            "parts": [{
                "text": text
            }]
        }]
    }

    # Header permintaan
    headers = {
        "Content-Type": "application/json"
    }

    # Melakukan permintaan POST
    response = requests.post(url, headers=headers, data=json.dumps(data))

    # Mengonversi respons ke JSON dan mencetaknya
    a = response.json()
    # print(a['candidates'][0]['content']['parts'][0]['text'])
    x = re.sub(r'\*', '', a['candidates'][0]['content']['parts'][0]['text'])
    return x


rObject = sr.Recognizer() 

def text_to_speech(text):
    tts = gTTS(text=text, lang='id')
    filename = "speech.mp3"
    tts.save(filename)
    os.system(f"ffplay -nodisp -autoexit -hide_banner {filename}")
    os.remove(filename)

while True:
    with sr.Microphone() as source: 
        print("Speak...")   
        rObject.adjust_for_ambient_noise(source, duration=1)
        audio = rObject.listen(source, timeout=0) 
        print("Stop.")
        try: 
            # text = rObject.recognize_google(audio, language ='en-US') 
            text = rObject.recognize_google(audio, language ='id') 
            print("You : "+ text)  
            print("Menunggu Jawaban...")  
            ans = getAnswer(text)
            text_to_speech(ans)
        except: 
            print("kesalahan")
