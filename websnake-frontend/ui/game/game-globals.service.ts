import { Injectable } from '@angular/core';
import * as PIXI from 'pixi.js';
import { Observable } from 'rxjs';

import { StoreService, CanvasReadyAction, ResourcesReadyAction } from 'websnake-frontend/store';

import { ALL_RESOURCES } from './resources';

@Injectable()
export class GameGlobalsService {
  private pixiApp: PIXI.Application;
  private pixiLoader: PIXI.loaders.Loader;

  constructor(private storeService: StoreService) {}

  canvasReady$ = this.storeService.selectLoad().getCanvasReady();
  resourcesReady$ = this.storeService.selectLoad().getResourcesReady();

  initializePixiApplication() {
    let pixiType = PIXI.utils.isWebGLSupported() ? 'WebGL' : 'canvas';
    PIXI.utils.sayHello(pixiType);

    this.pixiApp = new PIXI.Application();
    this.pixiApp.renderer.autoResize = true;
    this.pixiApp.renderer.resize(800, 600);
    this.storeService.dispatch(new CanvasReadyAction());
  }

  initializeResources() {
    this.pixiLoader = new PIXI.loaders.Loader();
    this.pixiLoader
      .add(ALL_RESOURCES)
      .load(() => this.storeService.dispatch(new ResourcesReadyAction()));
  }

  getApplication() {
    return Observable.combineLatest(this.canvasReady$, this.resourcesReady$)
      .filter(([cR, rR]) => cR && rR)
      .map(() => this.pixiApp);
  }

  getResource(resourceName: string) {
    return this.resourcesReady$.filter(x => x)
      .map(() => this.pixiLoader.resources[resourceName])
      .take(1);
  }
}
