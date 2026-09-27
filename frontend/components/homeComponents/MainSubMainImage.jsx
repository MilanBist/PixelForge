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
                                <h2>Transformations</h2>

                                <button onClick={()=>setSubmain("greyscale")}>GreyScale</button>
                                <button onClick={()=>setSubmain("increaseB")}>Increase Brightness</button>
                                <button onClick={()=>setSubmain("decreaseB")}>Decrease Brightness</button>
                                </>
                            )}

                            {main === "resize" && (
                                <>
                                <h2>Resize</h2>

                                <button onClick={()=>setSubmain("onex")}>1 x</button>
                                <button onClick={()=>setSubmain("twox")}>2 x</button>
                                <button onClick={()=>setSubmain("threex")}>3 x</button>
                                </>
                            )}

                            {main === "filters" && (
                                <>
                                <h2>Filters</h2>

                                <button onClick={()=>setSubmain("blur")}>Blur</button>
                                <button onClick={()=>setSubmain("sharpen")}>Sharpen</button>
                                <button onClick={()=>setSubmain("edge")}>Edge</button>
                                </>
                            )}
                        </div>
                    )}

                </div>
            </div>
        </>
    )
}