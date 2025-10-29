import json
#pip install googletrans
#pip install deep-translator
from deep_translator import GoogleTranslator

# Target language
language_target = 'de'
input_file = 'translation/en.json'
output_file = 'translation/'+language_target+'.json'
print(f"Translated JSON done : {output_file}")

# Initialize translator
translator = GoogleTranslator(source='en', target=language_target)

# Translate a text into a target language
def translate_json(text, target):
    # example target = 'fr'
    try:
        translation = translator.translate(text)
        return translation
    except Exception as e:
        print(f"Translation error : {e}")
        return text  # return text on error
    
with open(input_file, 'r', encoding='utf-8') as f:
    data = json.load(f)

for key, value in data.items():
    if "other" in value:
        value["other"] = translate_json(value["other"], language_target)

with open(output_file, 'w', encoding='utf-8') as f:
    json.dump(data, f, indent=4, ensure_ascii=False)

print(f"Translated JSON done : {output_file}")