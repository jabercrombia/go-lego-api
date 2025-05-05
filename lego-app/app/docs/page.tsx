// pages/docs.tsx
import React from "react";
import SwaggerComponent from "../../components/swaggerUI";

const DocsPage = () => {
  return (
    <div className="container mx-auto px-5 pt-5">
      <h1>API Documentation</h1>
      <SwaggerComponent />
    </div>
  );
};

export default DocsPage;
