import '../../styles/Home.css'
import MainSubMainImage from './MainSubMainImage'
import { useState } from 'react'
import axios from 'axios';

export default function AddImageCard({main, subMain, setMain, setSubMain, setOutput}){

    const[file, setFile] = useState(null);
    const handleFileChange = (evt)=>{
        setFile(evt.target.files[0]);
    }
    const getTransformedImage = async ()=>{
        if (file === null){
            alert("Upload the files first.");
            return;
        }

        if (main === null || subMain === null){
            alert("Set both of the main and submain properly first.");
            return;
        }

        const formData = new FormData();

        // set the field for the file
        formData.append("file", file);
        formData.append("mainTask", main);
        formData.append("subMainTask", subMain);


        axios.post("http://localhost:8081/api/transformImage", formData, {
            headers:{
                Authorization: `Bearer ${localStorage.getItem("accessToken")}`
            }
        }).then((resp)=>{
            console.log("Response is: ", resp);
            setOutput((prev)=>[
                ...prev, 
                resp.data["data"],
            ]);
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
                <button className="uploadraw-button" onClick={getTransformedImage} type='submit'>
                     Upload image file.
                </button>

                <div className="card-footer">
                    <span>Upload valid image file.</span>
                </div>
            </div>
        </>
    )
}