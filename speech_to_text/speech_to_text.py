import speech_recognition as sr
import logging 

# TODO: command line args to write output to a file or to a pipe 

# since this is the only python file we do logging for this file in here

logging.basicConfig(
    # level=logging.INFO,
    format="%(asctime)s [%(levelname)s] - %(message)s",
    handlers=[
        logging.FileHandler("./logging/app.log"),
        logging.StreamHandler()
    ]
)

"""
Currently writes the string to the given file
in the future might be used to send the strings over a linux pipe to the lexer
"""
def write_data(data):
    with open("user_output.txt", "a") as file:
        logging.INFO(f"writing {data} to file")
        file.write(data)
        file.close()
    return
    
"""
reads a continuious stream of data from the microphone 
"""
def read_from_microphone():
    r = sr.Recognizer() 
    while(1):    
        try:
            with sr.Microphone() as source:
                audio = r.listen(source)
                MyText = r.recognize_google(audio) # Using google to recognize audio
                print()
                print("Did you say ",MyText)
                print()
                write_data(MyText)
        except sr.UnknownValueError:
            print("google speech recognition couldnt understand audio")
        except sr.RequestError as e:
            print(f"Could not request results from google speech recognition; {e}")
        except KeyboardInterrupt:
            print("user interrupted the program, exiting")
            break
    return

def main():
    read_from_microphone()

# if __name__ == "__main__":
main()
