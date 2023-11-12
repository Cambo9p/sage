import speech_recognition as sr
import client


def write_data(data):
    with open("user_output.txt", "a") as file:
        print(f"data is {data}")
        file.write(data)
        file.close()
    return


def read_from_microphone():
    """
    reads a continuious stream of data from the microphone
    """
    r = sr.Recognizer()
    while 1:
        try:
            with sr.Microphone() as source:
                audio = r.listen(source)
                MyText = r.recognize_google(audio)
                print()
                print("Did you say ", MyText)
                print()
                # write_data(MyText)
                client.send_to_server(MyText)
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
