// pages/docs.tsx
import React from "react";
import SwaggerComponent from "../../components/swaggerUI";

const DocsPage = () => {
  return (
    <div className="container mx-auto">
      <h1>API Documentation</h1>
      <SwaggerComponent />
    </div>
  );
};

export default DocsPage;
