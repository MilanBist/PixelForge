import "../../styles/ImageButton.css"

export default function MainSubMainImage({main, subMain, setMain, setSubmain}){
    return(
        <>
            <div className="transformation-container">
                <div className="main-content">
                    <h3>Image transformation</h3>

                    <button className={main === "transformation" ? "selected" : ""} onClick={()=>setMain("transformation")}>
                        Transformation
                    </button>

                    <button className={main === "filters" ? "selected" : ""} onClick={()=>setMain("filters")}>
                        filters
                    </button>

                    <button className={main === "resize" ? "selected" : ""} onClick={()=>setMain("resize")}>
                        resize
                    </button>

                    {main && (
                        <div className="sub-section">
                            {main === "transformation" && (
                                <>
                                <h2>Transformations</h2>

                                <button className={subMain === "greyscale" ? "selected" : ""} onClick={()=>setSubmain("greyscale")}>GreyScale</button>
                                <button className={subMain === "increaseBrigtness" ? "selected" : ""} onClick={()=>setSubmain("increaseBrigtness")}>Increase Brightness</button>
                                <button className={subMain === "decreaseBrightness" ? "selected" : ""} onClick={()=>setSubmain("decreaseBrightness")}>Decrease Brightness</button>
                                </>
                            )}

                            {main === "resize" && (
                                <>
                                <h2>Resize</h2>

                                <button className={subMain === "onex" ? "selected" : ""} onClick={()=>setSubmain("onex")}>1 x</button>
                                <button className={subMain === "twox" ? "selected" : ""} onClick={()=>setSubmain("twox")}>2 x</button>
                                <button className={subMain === "threex" ? "selected" : ""} onClick={()=>setSubmain("threex")}>3 x</button>
                                </>
                            )}

                            {main === "filters" && (
                                <>
                                <h2>Filters</h2>

                                <button className={subMain === "blur" ? "selected" : ""} onClick={()=>setSubmain("blur")}>Blur</button>
                                <button className={subMain === "sharpen" ? "selected" : ""} onClick={()=>setSubmain("sharpen")}>Sharpen</button>
                                <button className={subMain === "edge" ? "selected" : ""} onClick={()=>setSubmain("edge")}>Edge</button>
                                </>
                            )}
                        </div>
                    )}

                </div>
            </div>
        </>
    )
}