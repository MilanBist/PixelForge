export default function MainSubMainImage({main, subMain, setMain, setSubmain}){
    return(
        <>
            <div className="transformation-container">
                <div className="main-content">
                    <h3>Image transformation</h3>

                    <button onClick={()=>setMain("transformation")}>
                        Transformation
                    </button>

                    <button onClick={()=>setMain("filters")}>
                        filters
                    </button>

                    <button onClick={()=>setMain("resize")}>
                        resize
                    </button>

                    {main && (
                        <div className="sub-section">
                            {main === "transformation" && (
                                <>
                                <h2>Transformation</h2>

                                <button onClick={()=>setSubmain("resize")}>Resize</button>
                                <button onClick={()=>setSubmain("rotate")}>Rotate</button>
                                <button onClick={()=>setSubmain("crop")}>Crop</button>
                                </>
                            )}

                            {main === "resize" && (
                                <>
                                <h2>Brightness</h2>

                                <button onClick={()=>setSubmain("brightness")}>Brightness</button>
                                <button onClick={()=>setSubmain("contrast")}></button>
                                <button onClick={()=>setSubmain("saturation")}>Saturation</button>
                                </>
                            )}

                            {main === "filters" && (
                                <>
                                <h2>Filters</h2>

                                <button onClick={()=>setSubmain("greyscale")}>Grayscale</button>
                                <button onClick={()=>setSubmain("blur")}>Blur</button>
                                <button onClick={()=>setSubmain("sharpen")}>Sharpen</button>
                                </>
                            )}
                        </div>
                    )}

                </div>
            </div>
        </>
    )
}