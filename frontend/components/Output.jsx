import ImageCard from "./imageComponents/ImageContainer";
import "../styles/Output.css"

function Output({output}){

    return (
        <>
            {output && output.map((data) => (
            <div className="input-file-section" key={data.inputFile.id}>

            <h2>{data.inputFile.name}</h2>
            <div className="image-grid">
                {data.images.map((image) => (
                    <ImageCard
                        key={image.id}
                        image={image}
                    />
                ))}
            </div>
        </div>
    ))}
        </>
    )
}
export default Output;