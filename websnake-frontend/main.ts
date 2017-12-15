import { enableProdMode } from '@angular/core';
import { platformBrowserDynamic } from '@angular/platform-browser-dynamic';
import { UIModule } from './ui/ui.module';

// depending on the env mode, enable prod mode or add debugging modules
if (['prod', 'deploy'].indexOf(process.env.ENV) > -1) {
  enableProdMode();
}

export function main() {
  return platformBrowserDynamic().bootstrapModule(UIModule);
}

if (document.readyState === 'complete') {
  main();
} else {
  document.addEventListener('DOMContentLoaded', main);
}
