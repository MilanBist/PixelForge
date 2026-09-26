import {useEffect } from "react";
import { getNewAccessToken } from "../utils/NewAccessToken";
import { useNavigate } from "react-router-dom";
import apiClient from "../utils/Base";
import "../styles/Uploaded.css"

function UploadedFiles({uploaded, setUploaded}){

    const navigate = useNavigate();
    
    
    async function getUploaded(){

        let accessToken = localStorage.getItem("accessToken");
        let refreshToken = localStorage.getItem("refreshToken");
        try{
            const resp = await apiClient.get("uploadedFiles", {
                headers:{
                    Authorization: `Bearer ${accessToken}`,
                }
            })
            console.log("Uploaded files are: ", resp.data.data);
            const data = resp.data.data;
            setUploaded(data);
            console.log("Uploaded files are: ",uploaded);
            console.log("Uploaded files are: ",data);
            return;
        }catch(err){
            const errorStatus = err.response?.status;
            if (errorStatus ===  401){
                // get new access token
                const response = await getNewAccessToken(refreshToken);
                console.log("Uploaded files:" ,response);
                if (response === 401){
                    // check for the refresh token availability
                     alert("You are logged out. Please login again");
                     setTimeout(()=>{
                        navigate("/login")
                     }, 1000);
                     return;
                }
                if (response === 200){
                    // successfully obtained new access token now you can again call the function of get history
                    getUploaded();
                    return 200;
                }
            }
            if (errorStatus === 400){
                alert('Wrong credentitals inputted.');
                return;
            }

            if (errorStatus === 500){
                alert("Internal server error. \n Try again later.");
                return;
            }
            return;
        }
        
    }

    useEffect(() => {
        let accessToken = localStorage.getItem("accessToken");
        let refreshToken = localStorage.getItem("refreshToken");
        if(accessToken === null || refreshToken === null){
            alert("You are not logged in. \n Redirecting to login page.");
            return;
        }
        getUploaded();
    }, []);

    useEffect(()=>{
        console.log("Type of uploaded: ", typeof uploaded);
        console.log(Array.isArray(uploaded));
    }, [uploaded]);

return (
    <div className="uploaded-files">
        {uploaded && uploaded.map((data) => (
            <div className="uploaded-file-card" key={data.uploadedId}>
                <div className="uploaded-file-icon">
                    RAW
                </div>

                <div className="uploaded-file-info">
                    <h3>{data.fileName}</h3>

                    <div className="uploaded-file-meta">
                        <span>{data.fileType}</span>
                        <span>•</span>
                        <span>{data.createdAt}</span>
                    </div>
                </div>

                {/* <button className="uploaded-file-button">
                    View
                </button> */}
            </div>
        ))}
    </div>
);
}
export default UploadedFiles;