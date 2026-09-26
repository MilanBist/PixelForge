import { useState } from "react";
import "../../styles/ImageCard.css";
import apiClient from "../../utils/Base";

export default function ImageCard({ image }) {

    const [preview, setPreview] = useState(null);

    async function handlePreview() {
        try {
            const response = await apiClient.get("/getSingleImage", {
                params: {
                    id: image.imageId,
                    mimetype: image.mimetype,
                },
                responseType: "blob", 
                headers: {
                    Authorization: `Bearer ${localStorage.getItem("accessToken")}`,
                }
            });

            // Convert received bytes into something <img> can display
            const imageUrl = URL.createObjectURL(response.data);
            console.log("Image url is: ",imageUrl);
            setPreview(imageUrl);

        } catch (error) {
            console.log("Failed to preview image:", error);
        }
    }

    function closePreview() {
        if (preview) {
            URL.revokeObjectURL(preview);
        }
        setPreview(null);
    }


    // handle download
    async function handledownload() {
        try {
            const response = await apiClient.get("/getSingleImage", {
                params: {
                    id: image.imageId,
                    mimetype: image.mimetype,
                },
                responseType: "blob", 
                headers: {
                    Authorization: `Bearer ${localStorage.getItem("accessToken")}`,
                }
            });

            // Convert received bytes into something <img> can display
            const downloadUrl = URL.createObjectURL(response.data);
            console.log("Image url is: ",downloadUrl);
            const link = document.createElement("a");
            link.href = downloadUrl;

            link.download = image.filename;

            document.body.appendChild(link);
            link.click();

            link.remove();
            URL.revokeObjectURL(downloadUrl);


        } catch (error) {
            console.log("Failed to preview image:", error);
        }
    }

    return (
        <>
            <div className="image-card">
                {/* Image */}
                <div className="image-container">
                    {/* Don't use storageKey here if it is only a backend path */}
                    <div className="image-placeholder">
                        Image
                    </div>
                </div>

                {/* Basic information */}
                <div className="image-details">
                    <h3>{image["imageName"]}</h3>
                    <div className="metadata">

                        <div>
                            <span>Type</span>
                            <strong>{image["mimetype"]}</strong>
                        </div>

                        <div>
                            <span>Size</span>
                            <strong>
                                {(image["filesize"] / 1024).toFixed(2)} KB
                            </strong>
                        </div>

                        <div>
                            <span>Dimensions</span>
                            <strong>
                                {image["width"]} × {image["height"]}
                            </strong>
                        </div>

                    </div>
                </div>

                {/* Actions */}
                <div className="image-actions">
                    <button onClick={handlePreview}>
                        Preview
                    </button>

                    <button onClick={handledownload}>
                        Download
                    </button>
                </div>

            </div>

            {/* Preview */}
            {preview && (
                <div className="preview-overlay">

                    <button
                        className="preview-close"
                        onClick={closePreview}
                    >
                        ✕
                    </button>

                    <img
                        src={preview}
                        alt={image["filename"]}
                        className="preview-image"
                    />

                    <button
                        className="preview-back"
                        onClick={closePreview}
                    >
                        ← Back
                    </button>

                </div>
            )}
        </>
    );
}