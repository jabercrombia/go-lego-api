declare module 'swagger-ui-react' {
    import { ComponentType } from 'react';
  
    export interface SwaggerUIProps {
      url?: string;
      spec?: object;
      dom_id?: string;
      onComplete?: () => void;
    }
  
    const SwaggerUI: ComponentType<SwaggerUIProps>;
    export default SwaggerUI;
  }
  