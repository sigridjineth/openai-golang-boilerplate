import argparse

from Chatbot import Chatbot


def main(access_token, prompt):
    chatbot = Chatbot(config={
        "access_token": access_token
    })

    response = ""

    for data in chatbot.ask(prompt):
        response = data["message"]

    print(response)

    return response


if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument("access_token", help="access token for the chatbot")
    parser.add_argument("prompt", help="prompt to pass to the chatbot")
    args = parser.parse_args()

    response = main(args.access_token, args.prompt)
