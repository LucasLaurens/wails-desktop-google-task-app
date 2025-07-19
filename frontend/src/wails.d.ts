export {}

declare global {
  interface Window {
    backend: {
      WailsEmitNewTaskCreationWithDescription(): Promise<string>;
    };
  }
}