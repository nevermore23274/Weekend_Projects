import os
import urllib.request
from llama_cpp import Llama
from huggingface_hub import snapshot_download

def download_file():
    filename = "llama-ar/llama-2-7b-langchain-chat-Q8_0.gguf"
    if not os.path.isfile(filename):
        model_id="YanaS/llama-2-7b-langchain-chat-GGUF"
        snapshot_download(repo_id=model_id, local_dir="llama-ar",
                          local_dir_use_symlinks=False, revision="main")
        print("file downloaded.")
    else:
        print("file exists already.")

def generate_text(
        prompt,
        max_tokens=256,
        temperature=1,
        top_p=0.5,
        echo=False,
        stop=None
):
    output = llm(
        prompt,
        max_tokens=max_tokens,
        temperature=temperature,
        top_p=top_p,
        echo=echo,
        stop=stop,
    )
    output_text = output["choices"][0]["text"].strip()
    log(output_text)

def log(txt):
    f = open(__file__ + '.log', "a")
    f.write(txt + '\r\n')
    f.close()

def generate_prompt_from_template(input):
    chat_prompt_template = f"""
    <|im_start|>user
    {input}<|im_end|>"""
    return chat_prompt_template

#ggml_model_path = "https://huggingface.co/YanaS/llama-2-7b-langchain-chat-GGUF"
#filename = "llama-ar/llama-2-7b-langchain-chat-Q8_0.gguf"
#download_file(ggml_model_path, filename)
download_file()

llm = Llama(model_path="llama-ar/llama-2-7b-langchain-chat-Q8_0.gguf", n_ctx=512, n_batch=126)

prompt = generate_prompt_from_template(
    "Translate the following from English to Arabic: Hello, what's your name?"
)

generate_text(
    prompt,
    max_tokens=356,
)
