import { Action } from '@ngrx/store';

export class CanvasReadyAction implements Action {
  static type = 'websnake/load/canvasReady';
  type = CanvasReadyAction.type;
}

export class ResourcesReadyAction implements Action {
  static type = 'websnake/load/resourcesReady';
  type = ResourcesReadyAction.type;
}
