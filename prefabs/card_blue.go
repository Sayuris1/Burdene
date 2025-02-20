components {
  id: "card_controller"
  component: "/scripts/card_controller.script"
  properties {
    id: "card_type"
    value: "1.0"
    type: PROPERTY_TYPE_NUMBER
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"card\"\n"
  "mask: \"cursor\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 77.387505\n"
  "  data: 124.59707\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"card\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 256.0\n"
  "  y: 256.0\n"
  "}\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/assets.atlas\"\n"
  "}\n"
  ""
}
embedded_components {
  id: "label"
  type: "label"
  data: "size {\n"
  "  x: 128.0\n"
  "  y: 32.0\n"
  "}\n"
  "color {\n"
  "  x: 0.0\n"
  "  y: 0.0\n"
  "  z: 0.0\n"
  "}\n"
  "text: \"Place a\\n"
  "\"\n"
  "  \"                  ball\\n"
  "\"\n"
  "  \"\\n"
  "\"\n"
  "  \"Ball Turns\\n"
  "\"\n"
  "  \"5               to a               \"\n"
  "font: \"/assets/Minecraft.font\"\n"
  "material: \"/builtins/fonts/label.material\"\n"
  ""
  position {
    z: 0.01
  }
}
embedded_components {
  id: "label1"
  type: "label"
  data: "size {\n"
  "  x: 128.0\n"
  "  y: 32.0\n"
  "}\n"
  "color {\n"
  "  x: 0.0\n"
  "}\n"
  "text: \"Blue\"\n"
  "font: \"/assets/Minecraft.font\"\n"
  "material: \"/builtins/fonts/label.material\"\n"
  ""
  position {
    x: 51.0
    y: -28.0
    z: 0.01
  }
}
embedded_components {
  id: "label2"
  type: "label"
  data: "size {\n"
  "  x: 128.0\n"
  "  y: 32.0\n"
  "}\n"
  "color {\n"
  "  x: 0.0\n"
  "}\n"
  "text: \"Blue\"\n"
  "font: \"/assets/Minecraft.font\"\n"
  "material: \"/builtins/fonts/label.material\"\n"
  ""
  position {
    x: -26.0
    y: 14.0
    z: 0.01
  }
}
embedded_components {
  id: "label3"
  type: "label"
  data: "size {\n"
  "  x: 128.0\n"
  "  y: 32.0\n"
  "}\n"
  "color {\n"
  "  y: 0.0\n"
  "  z: 0.0\n"
  "}\n"
  "text: \"Red\"\n"
  "font: \"/assets/Minecraft.font\"\n"
  "material: \"/builtins/fonts/label.material\"\n"
  ""
  position {
    x: -37.0
    y: -28.0
    z: 0.01
  }
}
