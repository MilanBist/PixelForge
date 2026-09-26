import axios from "axios";
import "../../styles/Home.css"
import apiClient from "../../utils/Base";
import { useState } from "react";
import { useNavigate } from "react-router-dom";


function RawExtractionCard({setOutputSection, setHistorySection, setUploadedSection}) {
  const [file, setFile] = useState(null);
  const navigate = useNavigate();

  // send the requet to get the new access token
  const getNewAccessToken = async (rt)=>{
      console.log("Refresh token is: ", rt);
      if (rt.length <1){
        return ["",-1];
      }
      console.log("reaching here.");
      const tokenBody = {
        refreshToken: rt,
      };

      console.log("The token body is: ", tokenBody);
      try{
        const resp = await apiClient.post("/refreshToken", tokenBody);
        let responseData = resp.data["message"];
        if(responseData ==="expired"){
          return ["expired", -2];
        }
        let accessToken = resp.data["data"]["accessToken"];
        console.log("The new access token is: ", accessToken);
        return [accessToken, 200];
      }catch(error){
        const status = error.response?.status;
        switch(status){
          case 500:
            let msg = error.response.data.message;
            console.log("From the new access token generating section. ", msg);
            return [error.response, 500];
          
          case 400:
            msg = error.response.data.message;
            console.log("From the new access token generating section. error: ", msg);
            alert(msg);
            return [error.response, 400];
        }
        console.log("Error in the generating new token section: ", error);
        return [error, status];
      } finally{
        console.log("Finished sending the response to backend.");
      }
    }

  const handleFileChange = (evt)=>{
    console.log("THis part is triggered.");
    setFile(evt.target.files[0]);
  }
  const decodeRawFile = async ()=>{
    if (file === null){
      alert("Upload the files first.");
      return;
    }
    console.log("reaching here.");

    const formData = new FormData();
    formData.append("file", file);


    const sendRequest = async (accessToken) => {
      return await axios.post("http://localhost:8081/api/getImages", formData, {
        headers: {
          "Content-Type":"multipart/form-data",
          Authorization: `Bearer ${accessToken}`,
        }
      });
    };
    
  try{
    const accessToken = localStorage.getItem("accessToken");
    let resp;
    try{
      resp = await sendRequest(accessToken);
      console.log("Initial response is: ", resp);
    } catch(err){
      const status = err.response.status;
      console.log(status);
      if (status !== 401){
        // send the throw
        throw err;
      }

      const refreshToken = localStorage.getItem("refreshToken");
      if (refreshToken == null){
        // redirect to the login page after 2 sec
        alert("You are not logged in: \n Please login.");
        alert("Redirecting you to login page.");
        setTimeout(()=>{
          navigate("/login")
        }, 1000);
        return;
      }
      // as if the refresh token is expired new login is required.
      const [token, code] = await getNewAccessToken(refreshToken);
      if (token === "expired"){
        alert("Please login again. \n Redirecting to the login page.")
        setTimeout(()=>{
          navigate("/login");
        },1000);

        return;
      }

      if (code !== 200){
        console.log("Error in gettting the error. Please login.");
        setTimeout(() => {
          navigate("/login");
        }, 3000);
      return;
      }


      localStorage.setItem("accessToken", token);
      resp = await sendRequest(token);
      }
      console.log("Responded data is: ", resp.data["data"]);
      // set the responded data in all of the parameters
      setOutputSection(prev =>[
        ...prev,
        resp.data["data"]
      ]);
      // setHistorySection(resp.data["data"]);
      // setUploadedSection(resp.data["data"]["inputFile"]);
    }catch(err){
      console.log("Request failed.", err);
    } finally{
      console.log("Finished sending requests.");
    }
  }
  return (
    <div className="asset-card">
      <div className="asset-card-header">
        <div className="asset-icon">☁</div>
      </div>
      <h3>RAW Extraction</h3>
      <p>
        Fetch processed raw camera buffers whose images you wanted.
      </p>

      <input type="file" className="uploadFile" onChange={handleFileChange}/>
      <button className="uploadraw-button" type="submit" onClick={decodeRawFile}>
        ▧ &nbsp; Upload Local RAW (card.raw)
      </button>

      <div className="card-footer">
        <span>Upload valid .raw file to get the response.</span>
      </div>
    </div>
  )
}

export default RawExtractionCard;