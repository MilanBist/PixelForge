import "../../styles/Home.css"


import { useRef } from "react";

export default function PromptSection() {
    const textareaRef = useRef(null);

    const handleInput = (e) => {
        const textarea = e.target;

        textarea.style.height = "auto";
        textarea.style.height = textarea.scrollHeight + "px";
    };

    return (
        <form className="search-container">

            <textarea
                ref={textareaRef}
                name="textarea"
                className="searcharea"
                placeholder="Describe what you want to create..."
                onInput={handleInput}
            />

            <button type="submit" className="search-button">
                <svg
                    xmlns="http://www.w3.org/2000/svg"
                    width="20"
                    height="20"
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke="currentColor"
                    strokeWidth="2"
                    strokeLinecap="round"
                    strokeLinejoin="round"
                >
                    <circle cx="11" cy="11" r="8"></circle>
                    <line x1="21" y1="21" x2="16.65" y2="16.65"></line>
                </svg>
            </button>

        </form>
    );
}