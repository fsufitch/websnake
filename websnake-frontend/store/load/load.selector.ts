import { Observable } from 'rxjs';

import { LoadState } from './load.state';

export class LoadSelector {
  constructor(private loadState$: Observable<LoadState>) {}

  getCanvasReady() {
    return this.loadState$.select(s => s.canvasReady);
  }

  getResourcesReady(){
    return this.loadState$.select(s => s.resourcesReady);
  }

}
