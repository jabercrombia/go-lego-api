"use client"

import React from "react";
import SwaggerUI from "swagger-ui-react";
import "swagger-ui-react/swagger-ui.css";

const SwaggerComponent = () => {
  return (
    <div>
      <SwaggerUI url="/docs/swagger.json" />
    </div>
  );
};

export default SwaggerComponent;
