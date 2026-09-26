import '../../styles/Home.css'
import MainSubMainImage from './MainSubMainImage'
import apiClient from '../../utils/Base'
import { useState } from 'react'

export default function AddImageCard({main, subMain, setMain, setSubMain}){

    const [file, setFile] = useState(null);

    const handleFileChange = (evt)=>{
        setFile(evt.target.files[0]);
    }
    const decodeRawFile = async ()=>{
        if (file === null){
            alert("Upload the files first.");
            return;
        }

        const formData = new FormData();
        // set the field for the file
        formData.append("file", file);
        apiClient.post("/getImages", formData).then((resp)=>{
            console.log("Response is: ", resp);
        }).catch((error)=>{
            const responseStatus = error.response.status;
            switch(responseStatus){
                case 400:
                    console.log(error.response.data.message);
                    alert(error.response.data.message);
                case 401:
                    let msg = error.response.data.message;
                    console.log(msg)
                    alert(msg);
                    // get the new token using the refresh token.
                    // if the refresh token still invalid then make the user log in again
                case 500:
                    msg = error.response.data.message;
                    console.log(msg)
                    alert(msg);           
                }
        }).finally(()=>{
            console.log("Image fetching completed.")
        })
    }
    return (
        <>
            <div className="asset-card">
                <div className="asset-card-header">
                    <div className="asset-icon">☁</div>
                </div>

                <h3>Image Transformation</h3>
                <p>
                    Add .jpg, .png file to transform the image.
                </p>

                <input type="file" className="uploadImage" onChange={handleFileChange}/>

                <MainSubMainImage main={main} subMain={subMain} setMain={setMain} setSubmain={setSubMain}/>
                <button className="uploadraw-button" onSubmit={decodeRawFile} type='submit'>
                    ▧ &nbsp; Upload image file.
                </button>

                <div className="card-footer">
                    <span>Upload valid image file.</span>
                </div>


                {/* <PromptSection/> */}
            </div>
        </>
    )
}