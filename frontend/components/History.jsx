import { useEffect } from "react";
import apiClient from "../utils/Base";
import { useNavigate } from "react-router-dom";
import { getNewAccessToken } from "../utils/NewAccessToken";
import ImageCard from "./imageComponents/ImageContainer";


function History({output, history, setHistory}){
    const navigate = useNavigate();
    let accessToken = localStorage.getItem("accessToken");
    let refreshToken = localStorage.getItem("refreshToken");
    async function getHistory(){
        try{
            const resp = await apiClient.get("history", {
                headers:{
                    Authorization: `Bearer ${accessToken}`,
                }
            })

            // if successfull then just return
            console.log(resp.status);
            if (resp.status === 200){
                setHistory(resp.data["data"]);
                console.log("History is: ",history);
                return;
            }
        }catch(err){
            const errorStatus = err.response?.status;
            if (errorStatus ===  401){
                // get new access token
                const response = await getNewAccessToken(refreshToken);
                console.log("Historical Data:" ,response);
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
                    getHistory();
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
        let refreshToken = localStorage.getItem("refreshToken");
        let accessToken = localStorage.getItem("accessToken");

        if(accessToken === null || refreshToken === null){
            alert("You are not logged in. \n Redirecting to login page.");
            return;
        }


        getHistory();
    }, [output]);

    return (
        <>
            <div className="history-page">
                     {history && history.map((data) => (
                            <div className="input-file-section" key={data.inputFile.uploadedId}>
                            <h2>{data.inputFile.name}</h2>  
                            <h3>{data.inputFile.createdAt}</h3>
                            <div className="image-grid">
                                {data.images.map((image) => (
                                    <ImageCard
                                        key={image.imageId}
                                        image={image}
                                    />
                                ))}
                            </div>
                        </div>
                    ))}
            </div>

        </>
    )
}

export default History;