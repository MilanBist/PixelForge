import apiClient from "./Base";

export async function getNewAccessToken(refreshToken){
    console.log("Refresh token is: ", refreshToken);
    let tokenBody = {
        refreshToken: refreshToken,
    }
    try{
        const resp = await apiClient.post("/refreshToken", tokenBody);
        localStorage.setItem("accessToken", resp.data["data"]["accessToken"]);
        return 200;
    }catch(err){
        return err.response.status;
    }
}