import "../styles/Home.css";
import Navbar from "./homeComponents/Navbar";
import RawExtractionCard from "./homeComponents/RawExtraction";
import AddImageCard from "./homeComponents/ImageuploadSection";
import PromptSection from "./homeComponents/PromptSection";
import { useState } from "react";


function SourceAssets({setOutputSection, setHistorySection, setUploadedSection}) {
    // set the states of the main and subsections here
    const[main, setMain] = useState(null);
    const[subMain, setSubMain] = useState(null);


  return (
    <section className="section">
      <div className="section-heading">
        <h2>Source Assets</h2>
        <span>
          STAGE 01 / INPUT
        </span>
      </div>
      <div className="asset-grid">
        <RawExtractionCard setOutputSection={setOutputSection} setHistorySection={setHistorySection} setUploadedSection={setUploadedSection}/>
        <AddImageCard main = {main}  subMain = {subMain}  setMain={setMain} setSubMain={setSubMain}/>
      </div>
    </section>
  );
}

export default function Home({setOutputSection, setHistorySection, setUploadedSection}) {
  return (
    <main className="home-page">
      <SourceAssets setOutputSection={setOutputSection} setHistorySection={setHistorySection} setUploadedSection={setUploadedSection}/>
      <p className="note">Use the prompt bar when you upload the image file only.(Optional)</p>
      <PromptSection />
    </main>
  );
}