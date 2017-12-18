import { Component, ElementRef, OnInit } from '@angular/core';
import * as PIXI from 'pixi.js';

import { GameGlobalsService } from './game-globals.service';

import { CAT } from './resources';

@Component({
  selector: 'game-canvas',
  template: require('./game-canvas.component.html'),
})
export class GameCanvasComponent implements OnInit {
  private myElement: HTMLElement = this.elementRef.nativeElement;

  application$ = this.gameGlobalsService.getApplication().take(1);
  catSprite$ = this.gameGlobalsService.getResource(CAT)
    .map(r => new PIXI.Sprite(r.texture));

  constructor(
    private elementRef: ElementRef,
    private gameGlobalsService: GameGlobalsService,
  ) {}


  ngOnInit() {
    this.gameGlobalsService.getApplication().take(1)
      .subscribe(app => this.myElement.appendChild(app.view));

    this.application$.zip(this.catSprite$).subscribe(([app, sprite]) => {
      app.stage.addChild(sprite);
    });
  }
}
