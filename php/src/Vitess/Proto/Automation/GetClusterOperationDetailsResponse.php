<?php
// DO NOT EDIT! Generated by Protobuf-PHP protoc plugin 1.0
// Source: automation.proto
//   Date: 2016-01-22 01:34:21

namespace Vitess\Proto\Automation {

  class GetClusterOperationDetailsResponse extends \DrSlump\Protobuf\Message {

    /**  @var \Vitess\Proto\Automation\ClusterOperation */
    public $cluster_op = null;
    

    /** @var \Closure[] */
    protected static $__extensions = array();

    public static function descriptor()
    {
      $descriptor = new \DrSlump\Protobuf\Descriptor(__CLASS__, 'automation.GetClusterOperationDetailsResponse');

      // OPTIONAL MESSAGE cluster_op = 2
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 2;
      $f->name      = "cluster_op";
      $f->type      = \DrSlump\Protobuf::TYPE_MESSAGE;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $f->reference = '\Vitess\Proto\Automation\ClusterOperation';
      $descriptor->addField($f);

      foreach (self::$__extensions as $cb) {
        $descriptor->addField($cb(), true);
      }

      return $descriptor;
    }

    /**
     * Check if <cluster_op> has a value
     *
     * @return boolean
     */
    public function hasClusterOp(){
      return $this->_has(2);
    }
    
    /**
     * Clear <cluster_op> value
     *
     * @return \Vitess\Proto\Automation\GetClusterOperationDetailsResponse
     */
    public function clearClusterOp(){
      return $this->_clear(2);
    }
    
    /**
     * Get <cluster_op> value
     *
     * @return \Vitess\Proto\Automation\ClusterOperation
     */
    public function getClusterOp(){
      return $this->_get(2);
    }
    
    /**
     * Set <cluster_op> value
     *
     * @param \Vitess\Proto\Automation\ClusterOperation $value
     * @return \Vitess\Proto\Automation\GetClusterOperationDetailsResponse
     */
    public function setClusterOp(\Vitess\Proto\Automation\ClusterOperation $value){
      return $this->_set(2, $value);
    }
  }
}
