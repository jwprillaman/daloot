import {Component, OnInit} from 'angular2/core';
import {ANGULAR2_GOOGLE_MAPS_DIRECTIVES} from 'angular2-google-maps/core';

@Component({
  selector: 'my-map',
  directives: [ANGULAR2_GOOGLE_MAPS_DIRECTIVES], // this loads all angular2-google-maps directives in this component
  // the following line sets the height of the map - Important: if you don't set a height, you won't see a map!!
  styles: [`
    .sebm-google-map-container {
      height: 300px;
    }
  `],
  template: `
    <h1>My first angular-google-maps app!</h1>

    <!-- this creates a google map on the page with the given lat/lng from the component as the initial center of the map: -->

    <sebm-google-map [latitude]="lat" [longitude]="lng">
    </sebm-google-map>
  `
})
export class MapComponent {
  lat: number = 51.678418;
  lng: number = 7.809007;
}
