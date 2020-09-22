<?php

namespace m{module}\controller;

use think\exception\ValidateException;
use Cgf\Framework\Thinkphp\BaseController;
class {moduleUpper} extends BaseController
{
    function initialize(){
        parent::initialize();
        $this->m = new \{module}\model\{moduleUpper}();
    }
    function getModelDir(){
        return "\\{module}\\model";
    }

}
