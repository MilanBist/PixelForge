import '../../styles/Home.css'
import MainSubMainImage from './MainSubMainImage'
import { useState } from 'react'
import axios from 'axios';
import { getNewAccessToken } from '../../utils/NewAccessToken';

export default function AddImageCard({main, subMain, setMain, setSubMain, setOutput}){

    const[file, setFile] = useState(null);
    const handleFileChange = (evt)=>{
        setFile(evt.target.files[0]);
    }
    const getTransformedImage = async () => {
        let refreshToken = localStorage.getItem("refreshToken");
        if (file === null) {
            alert("Upload the files first.");
            return;
        }

        if (main === null || subMain === null) {
            alert("Set both of the main and submain properly first.");
            return;
        }

        const formData = new FormData();

        // Set the fields
        formData.append("file", file);
        formData.append("mainTask", main);
        formData.append("subMainTask", subMain);

        try {
            const resp = await axios.post(
                "http://localhost:8081/api/transformImage",formData,{
                    headers: {
                        Authorization: `Bearer ${localStorage.getItem("accessToken")}`
                    }
                }
            );

            console.log("Response is: ", resp);

            setOutput((prev) => [
                ...prev,
                resp.data["data"]
            ]);
        } catch (error) {
            const responseStatus = error.response?.status;
            switch (responseStatus) {
                case 400: {
                    console.log(error.response.data.message);
                    alert(error.response.data.message);
                    break;
                }

                case 401: {
                    const msg = error.response.data.message;
                    const response = await getNewAccessToken(refreshToken);
                    console.log("Historical Data:", response);
                    if (response === 401) {
                        alert("You are logged out. Please login again");
                        setTimeout(() => {
                            navigate("/login");
                        }, 1000);

                        return;
                    }

                    if (response === 200) {
                        return getTransformedImage();
                    }
                    alert(msg);
                    break;
                }
                case 500: {
                    const msg = error.response.data.message;
                    console.log(msg);
                    alert(msg);
                    break;
                }
                default: {
                    console.log("Unexpected error:", error);
                    alert("Something went wrong.");
                    break;
                }
            }

        } finally {
            console.log("Image fetching completed.");
        }
    };
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