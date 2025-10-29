import json
import os

# Adapt format of i18n for fyne lang
def transform_json(file_path):
    with open(file_path, 'r', encoding='utf-8') as file:
        data = json.load(file)
    
    transformed_data = {key: value["other"] for key, value in data.items()}

    with open(file_path, 'w', encoding='utf-8') as file:
        json.dump(transformed_data, file, indent=4)

folder_path = "translation"

# Apply for all .json
for filename in os.listdir(folder_path):
    if filename.endswith(".json"):
        transform_json(os.path.join(folder_path, filename))

print("finished")