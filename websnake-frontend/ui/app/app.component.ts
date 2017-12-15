import { Component } from '@angular/core';

@Component({
  selector: 'ng2app',
  template: require('./app.component.html'),
})
export class AppComponent {
  apiHost: string = (<any>window).WEBSNAKE_API_HOST;
}
