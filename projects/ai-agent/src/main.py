import os
import sys

from dotenv import load_dotenv
from google import genai
from google.genai import types

from config import MAX_ITER, system_prompt
from functions.call_function import available_functions, call_function


def main():
    load_dotenv()
    api_key = os.environ.get("GEMINI_API_KEY")
    client = genai.Client(api_key=api_key)

    print(sys.argv)
    if len(sys.argv) < 2:
        print("Error: no prompt was provided")
        sys.exit(1)
    user_prompt = sys.argv[1]

    messages = [
        types.Content(role="user", parts=[types.Part(text=user_prompt)]),
    ]

    verbosity = False
    if "--verbose" in sys.argv:
        verbosity = True
        print(f"User prompt: {user_prompt}")

    iter = 0
    while True:
        iter += 1
        if iter > MAX_ITER:
            print(f"Maximum iterations ({MAX_ITER}) reached.")
            sys.exit(1)

        try:
            final_response = generate_content(client, messages, verbosity)
            if final_response:
                print("Final response:")
                print(final_response)
                break
        except Exception as e:
            print(f"Encountered an exception: {e}")


def generate_content(client, messages, verbose):
    response = client.models.generate_content(
        model="gemini-2.0-flash-001",
        contents=messages,
        config=types.GenerateContentConfig(
            tools=[available_functions], system_instruction=system_prompt
        ),
    )

    for candidate in response.candidates:
        messages.append(candidate.content)

    if verbose:
        print(f"Prompt tokens: {response.usage_metadata.prompt_token_count}")
        print(f"Response tokens: {response.usage_metadata.candidates_token_count}")
    if not response.function_calls:
        return response.text

    tool_response_parts = []
    for function_call_part in response.function_calls:
        print(f"Calling function: {function_call_part.name}({function_call_part.args})")
        fn_response = call_function(function_call_part, verbose)
        if not fn_response.parts[0].function_response.response:
            raise Exception("Fatal: fucntion call error")

        tool_response_parts.append(
            types.Part.from_function_response(
                name=function_call_part.name,
                response={"result": fn_response},
            )
        )

    messages.append(types.Content(role="tool", parts=tool_response_parts))


if __name__ == "__main__":
    main()
