from fastapi import FastAPI, UploadFile, File
from typing import Optional
from fastapi.responses import FileResponse
from PIL import Image
import uuid
import ffmpeg
#import magic
from mimetypes import MimeTypes
from fastapi.exceptions import HTTPException
import urllib.request
import string
import av
import subprocess
import os

def extract_frames(video_path, output_folder):
    if not os.path.exists(output_folder):
        os.makedirs(output_folder)
    
    command = [
        'ffmpeg',
        '-ss', '00:00:00', '-i', video_path,              # Input file
        '-frames:v', '1',    # Video filter to set frame rate
        os.path.join(output_folder, 'frame_%04d.png')  # Output file pattern
    ]
    
    subprocess.run(command, check=True)

app = FastAPI()

@app.put("/api/files")
def upload_file(file: UploadFile):
    if file.content_type.partition("/")[0] != "image" and file.content_type.partition("/")[0] != "video":
        return {"message": "Invalid document type"}
        #raise HTTPException(400, detail="Invalid document type")
    #return FileResponse(file, filename=file.filename, media_type=file.content_type)
    try:
        file_id = uuid.uuid4()
        file_path = f"/Users/ramilbagaviev/Downloads/Python/TestFiles/{file.filename}"
        with open(file_path, "wb") as f:
            f.write(file.file.read())
        with open(f"/Users/ramilbagaviev/Downloads/Python/file_id_name.txt", "a") as ffile:
            ffile.write(str(file_id) + " " + file.filename + "\n")
    except Exception as e:
        return {"message": e.args}
    return {"file_id": file_id, "filename": file.filename, "mime": file.content_type, "file_size": file.file._max_size}

@app.put("/api/files/{uuid}") #/api/files/{uuid}?width&height
async def read_id(s_id: str, a: Optional[int] = None, b: Optional[int] = None):
    with open(f"/Users/ramilbagaviev/Downloads/Python/file_id_name.txt", "r") as file:
        line = file.readline()
        while line:
            if line.partition(" ")[0] == s_id:
                nameOfFile = line.partition(" ")[2].rstrip()
                #mime = magic.Magic(mime=True)
                #filetype = magic.from_file(f"/Users/ramilbagaviev/Downloads/Python/TestFiles/{nameOfFile}").partition("/")[0]
                mime = MimeTypes()
                url = urllib.request.pathname2url(f"/Users/ramilbagaviev/Downloads/Python/TestFiles/{nameOfFile}")
                filetype = mime.guess_type(url)[0].partition("/")[0]
                #return filetype
                if filetype == "image":
                    if a == None or b == None:
                        return {"upd_file": Image.open(f"/Users/ramilbagaviev/Downloads/Python/TestFiles/{nameOfFile}").show()}
                    return {"upd_file": Image.open(f"/Users/ramilbagaviev/Downloads/Python/TestFiles/{nameOfFile}").resize((a,b)).show()}
                #return FileResponse(path=f"/Users/ramilbagaviev/Downloads/Python/TestFiles/{nameOfFile}", filename=nameOfFile, media_type='multipart/form-data')
                #return {"filename": line.partition(" ")[2].rstrip()}
                if filetype == "video":
                    extract_frames(f"/Users/ramilbagaviev/Downloads/Python/TestFiles/{nameOfFile}", f"/Users/ramilbagaviev/Downloads/Python/TestFiles")
                    return {"first_frame": Image.open(f"/Users/ramilbagaviev/Downloads/Python/TestFiles/frame_0001.png").show()}
                    return av.open(f"/Users/ramilbagaviev/Downloads/Python/TestFiles/{nameOfFile}").streams.video[0].frames
            line = file.readline()
    raise HTTPException(404, detail="No such file")
    #return {"message": "No such file"}  
